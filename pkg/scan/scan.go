package scan

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/url"
	"time"

	"golang.org/x/net/proxy"
)

// Scanner manages TCP scanning.
type Scanner struct {
	Port        int
	Request     []byte
	Threads     int
	DialTimeout time.Duration
	ReadTimeout time.Duration
	Blacklist   *Blacklist
	Dialer      proxy.ContextDialer
}

// NewScanner returns new Scanner with threads count.
func NewScanner(threads int) *Scanner {
	return &Scanner{
		Port:        80,
		Request:     []byte("GET / HTTP/1.0"),
		Threads:     threads,
		DialTimeout: 1 * time.Second,
		ReadTimeout: 5 * time.Second,
		Blacklist:   &Blacklist{},
		Dialer:      &net.Dialer{},
	}
}

// Read header from addr as bytes
func (s *Scanner) scan(addr string) ([]byte, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), s.DialTimeout)
	defer cancel()
	conn, err := s.Dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, false
	}
	defer conn.Close()
	buf := make([]byte, 4096)
	conn.SetDeadline(time.Now().Add(s.ReadTimeout))
	_, err = conn.Write(s.Request)
	if err != nil {
		return nil, false
	}
	_, err = conn.Write([]byte("\r\n\r\n"))
	if err != nil {
		return nil, false
	}
	n, err := conn.Read(buf)
	if err != nil {
		return nil, false
	}
	data := buf[:n]
	parts := bytes.SplitN(data, []byte("\r\n\r\n"), 2)
	if len(parts) != 2 {
		return nil, false
	}
	return parts[0], true
}

// Infinitely scan random addresses and pump valid results to ch
func (s *Scanner) worker(ch chan *Result) {
	for {
		var ip net.IP
		for {
			ip = randomIP()
			if !s.Blacklist.Contains(ip) {
				break
			}
		}
		port := s.Port
		addr := fmt.Sprintf("%v:%d", ip, port)
		headers, ok := s.scan(addr)
		if ok {
			ch <- &Result{IP: ip, Port: port, Headers: string(headers)}
		}
	}
}

// SetProxy sets a proxy for the scanner by URL.
func (s *Scanner) SetProxy(proxyURL string) error {
	u, err := url.Parse(proxyURL)
	if err != nil {
		return err
	}
	d, err := proxy.FromURL(u, proxy.Direct)
	if err != nil {
		return err
	}
	cd, ok := d.(proxy.ContextDialer)
	if !ok {
		return errors.New("proxy doesn't implement ContextDialer")
	}
	s.Dialer = cd
	return nil
}

// Start workers concurrently and return result channel.
func (s *Scanner) Start() chan *Result {
	ch := make(chan *Result)
	for i := 0; i < s.Threads; i++ {
		go s.worker(ch)
	}
	return ch
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
