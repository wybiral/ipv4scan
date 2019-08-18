package scan

import "net"

// Result represents one scan result.
type Result struct {
	IP      net.IP `json:"ip"`
	Port    int    `json:"port"`
	Headers string `json:"headers"`
}
