package scan

import (
	"bufio"
	"net"
	"os"
	"strings"
)

// Blacklist manages blacklisted IP network.
type Blacklist struct {
	networks []*net.IPNet
}

// Contains returns true/false if blacklist contains IP.
func (b *Blacklist) Contains(ip net.IP) bool {
	for _, ipnet := range b.networks {
		if ipnet.Contains(ip) {
			return true
		}
	}
	return false
}

// Add adds a network to blacklist from CIDR notation.
func (b *Blacklist) Add(cidr string) error {
	_, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return err
	}
	b.networks = append(b.networks, ipnet)
	return nil
}

// Parse parses file containing lines of CIDR notation into blacklist.
func (b *Blacklist) Parse(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		parts := strings.SplitN(scan.Text(), "#", 2)
		cidr := strings.Trim(parts[0], " \t")
		if len(cidr) == 0 {
			continue
		}
		err := b.Add(cidr)
		if err != nil {
			return err
		}
	}
	return scan.Err()
}
