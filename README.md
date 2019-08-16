# ipv4scan
IoT device scanner.
### Install dependencies
```
go get github.com/wybiral/ipv4scan
```
### Build
```
go build github.com/wybiral/ipv4scan/cmd/scan
```
### Usage
Specify thread count:
```
./scan -n 1000
```
To start scanning with a blacklist:
```
./scan -b blacklist.conf
```
To start scanning through a SOCKS proxy (such as Tor):
```
./scan -p socks5://127.0.0.1:9050
```
Example blacklist file located at [configs/blacklist.conf](https://github.com/wybiral/ipv4scan/blob/master/configs/blacklist.conf)
