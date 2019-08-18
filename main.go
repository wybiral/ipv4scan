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
	// Setup blacklist flag
	blacklist := "blacklist.conf"
	flag.StringVar(
		&blacklist,
		"b",
		blacklist,
		"blacklist file",
	)
	// Setup threads flag
	threads := 100
	flag.IntVar(
		&threads,
		"n",
		threads,
		"number of scanner threads",
	)
	// Setup proxy flag
	proxyURL := ""
	flag.StringVar(
		&proxyURL,
		"p",
		proxyURL,
		"proxy URL",
	)
	flag.Parse()
	scanner := scan.NewScanner(threads)
	// optionally setup blacklist
	if len(blacklist) > 0 {
		err := scanner.Blacklist.Parse(blacklist)
		if err != nil {
			log.Fatal(err)
		}
	}
	// optionally setup proxy
	if len(proxyURL) > 0 {
		err := scanner.SetProxy(proxyURL)
		if err != nil {
			log.Fatal(err)
		}
	}
	encoder := json.NewEncoder(os.Stdout)
	for result := range scanner.Start() {
		err := encoder.Encode(result)
		if err != nil {
			return
		}
	}
}
