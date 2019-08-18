# ipv4scan
IoT device scanner.
### Install dependencies
```
go get github.com/wybiral/ipv4scan
```
### Build
```
go build github.com/wybiral/ipv4scan
```
### Usage
Specify thread count:
```
./ipv4scan -n 1000
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
