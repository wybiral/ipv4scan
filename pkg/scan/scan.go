package scan

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/wybiral/ipv4scan/pkg/types"
	"golang.org/x/net/proxy"
)

// Scanner manages TCP scanning.
type Scanner struct {
	Threads     int
	DialTimeout time.Duration
	ReadTimeout time.Duration
	Blacklist   *Blacklist
	Dialer      proxy.ContextDialer
}

// NewScanner returns new Scanner with threads count.
func NewScanner(threads int) *Scanner {
	return &Scanner{
		Threads:     threads,
		DialTimeout: 1 * time.Second,
		ReadTimeout: 5 * time.Second,
		Blacklist:   &Blacklist{},
		Dialer: &net.Dialer{
			Timeout: 5 * time.Second,
		},
	}
}

// Return random IPv4 address
func randomIP() net.IP {
	i := rand.Uint32()
	a := byte(i)
	b := byte(i >> 8)
	c := byte(i >> 16)
	d := byte(i >> 24)
	return net.IPv4(a, b, c, d)
}

// Read header from addr as bytes
func (s *Scanner) scan(addr string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.DialTimeout)
	defer cancel()
	conn, err := s.Dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	buf := make([]byte, 4096)
	conn.SetDeadline(time.Now().Add(s.ReadTimeout))
	_, err = conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))
	if err != nil {
		return nil, err
	}
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	data := buf[:n]
	parts := bytes.SplitN(data, []byte("\r\n\r\n"), 2)
	if len(parts) == 2 {
		return parts[0], nil
	}
	return nil, errors.New("scanner: invalid http header")
}

// Infinitely scan random addresses and pump valid results to ch
func (s *Scanner) worker(ch chan *types.Result) {
	for {
		var ip net.IP
		for {
			ip = randomIP()
			if !s.Blacklist.Contains(ip) {
				break
			}
		}
		port := 80
		addr := fmt.Sprintf("%v:%d", ip, port)
		headers, err := s.scan(addr)
		if err != nil {
			continue
		}
		ch <- &types.Result{IP: ip, Port: port, Headers: string(headers)}
	}
}

// Start a number of workers concurrently and return result channel
func (s *Scanner) Start() chan *types.Result {
	ch := make(chan *types.Result)
	for i := 0; i < s.Threads; i++ {
		go s.worker(ch)
	}
	return ch
}
