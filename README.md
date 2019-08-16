# scan
IoT device scanner.
### Install dependencies
```
go get github.com/wybiral/scan
```
### Build
```
go build github.com/wybiral/scan/cmd/scan
```
### Usage
To start scanning with a blacklist:
```
./scan -b blacklist.conf
```
To start scanning through a SOCKS proxy (such as Tor):
```
./scan -p socks5://127.0.0.1:9050
```
Example blacklist file located at [configs/blacklist.conf](https://github.com/wybiral/scan/blob/master/configs/blacklist.conf)
