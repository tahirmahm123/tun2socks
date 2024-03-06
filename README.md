# Client for Tun2Socks
#### This project uses Tun2Sock Library from [Tun2Sock Core](github.com/tahirmahm123/go-tun2socks) which is actually a [fork](github.com/benchibench/go-tun2socks) but original code exists only on [Go Packages](https://pkg.go.dev/github.com/benchibench/go-tun2socks)

#### This repo is a simple client to use tun2socks.

### For Compilation
`
    make
`
### Run Tun2Socks with admin privileges
`
./build/tun2socks -tunName tun1 -tunAddr 240.0.0.2 -tunGw 240.0.0.1 -proxyType shadowsocks -proxyServer "1.2.3.4:8388" -loglevel "debug" -proxyPassword "xxxxxxx"
`

Assuming you have 1.2.3.4 your Proxy Server and 192.168.1.1 as Local Gateway
### Add Routes
`
route delete default
route add default 240.0.0.1
route add 1.2.3.4/32 192.168.1.1
`
### After Exit Make sure to remove routes
`
route delete default
route delete 240.0.0.1
route delete 1.2.3.4/32 192.168.1.1
route add default 192.168.1.1
`