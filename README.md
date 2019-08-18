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
To start scanning with a blacklist:
```
./ipv4scan -b blacklist.conf
```
To start scanning through a SOCKS proxy (such as Tor):
```
./ipv4scan -p socks5://127.0.0.1:9050
```
Example blacklist file: [blacklist.conf](https://github.com/wybiral/ipv4scan/blob/master/blacklist.conf)
