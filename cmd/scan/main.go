package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/url"
	"os"
	"time"

	"github.com/wybiral/ipv4scan/pkg/scan"
	"golang.org/x/net/proxy"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	// Setup blacklist flag
	blacklist := "configs/blacklist.conf"
	flag.StringVar(
		&blacklist,
		"b",
		blacklist,
		"blacklist file containing CIDR notation",
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
		"proxy URL (optional)",
	)
	flag.Parse()
	// Create scanner and (optionally) load blacklist
	scanner := scan.NewScanner(threads)
	if len(blacklist) > 0 {
		err := scanner.Blacklist.Parse(blacklist)
		if err != nil {
			log.Fatal(err)
		}
	}
	// optionally setup proxy
	if len(proxyURL) > 0 {
		u, err := url.Parse(proxyURL)
		if err != nil {
			log.Fatal(err)
		}
		d, err := proxy.FromURL(u, proxy.Direct)
		if err != nil {
			log.Fatal(err)
		}
		cd, ok := d.(proxy.ContextDialer)
		if !ok {
			log.Fatal("proxy doesn't implement ContextDialer")
		}
		scanner.Dialer = cd
	}
	encoder := json.NewEncoder(os.Stdout)
	for result := range scanner.Start() {
		err := encoder.Encode(result)
		if err != nil {
			return
		}
	}
}
