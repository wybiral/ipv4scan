# ipv4scan
IoT device scanner. Randomly scan the IPv4 address space and collect a real-time JSON stream of all HTTP devices that are found. This stream can then by piped into real-time filters and analytics or piped into storage to be searched as your own personal shodan.

## Download latest
All binary releases can be found [here](https://github.com/wybiral/ipv4scan/releases).

## Build from source
```
go get github.com/wybiral/ipv4scan
go build github.com/wybiral/ipv4scan
```

## Usage
Specify thread count:
```
./ipv4scan -n 1000
```
Specify HTTP port:
```
./ipv4scan -p 8080
```
Specify blacklist:
```
./ipv4scan -b blacklist.conf
```
Scan through SOCKS proxy (such as Tor):
```
./ipv4scan -x socks5://127.0.0.1:9050
```
Specify HTTP request line:
```
./ipv4scan -r "GET /some_resource HTTP/1.1"
```
Example blacklist file: [blacklist.conf](https://github.com/wybiral/ipv4scan/blob/master/blacklist.conf)
