package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/wybiral/ipv4scan/pkg/scan"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// setup blacklist flag
	blacklist := "blacklist.conf"
	flag.StringVar(
		&blacklist,
		"b",
		blacklist,
		"blacklist file",
	)
	// setup threads flag
	threads := 100
	flag.IntVar(
		&threads,
		"n",
		threads,
		"number of scanner threads",
	)
	// setup proxy flag
	proxyURL := ""
	flag.StringVar(
		&proxyURL,
		"p",
		proxyURL,
		"proxy URL",
	)
	// setup request flag
	request := "GET / HTTP/1.0"
	flag.StringVar(
		&request,
		"r",
		request,
		"HTTP request",
	)
	flag.Parse()
	scanner := scan.NewScanner(threads)
	// setup blacklist
	if len(blacklist) > 0 {
		err := scanner.Blacklist.Parse(blacklist)
		if err != nil {
			log.Fatal(err)
		}
	}
	// setup proxy
	if len(proxyURL) > 0 {
		err := scanner.SetProxy(proxyURL)
		if err != nil {
			log.Fatal(err)
		}
	}
	// setup request
	if len(request) > 0 {
		scanner.Request = []byte(request)
	}
	encoder := json.NewEncoder(os.Stdout)
	for result := range scanner.Start() {
		err := encoder.Encode(result)
		if err != nil {
			return
		}
	}
}
