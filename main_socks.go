//go:build socks
// +build socks

package main

import (
	"net"

	"github.com/tahirmahm123/go-tun2socks/common/log"
	"github.com/tahirmahm123/go-tun2socks/core"
	"github.com/tahirmahm123/go-tun2socks/proxy/socks"
)

func init() {
	args.addFlag(fProxyServer)
	args.addFlag(fUdpTimeout)
	args.addFlag(fApplog)

	registerHandlerCreater("socks", func() {
		// Verify proxy server address.
		proxyAddr, err := net.ResolveTCPAddr("tcp", *args.ProxyServer)
		if err != nil {
			log.Fatalf("invalid proxy server address: %v", err)
		}
		proxyHost := proxyAddr.IP.String()
		proxyPort := uint16(proxyAddr.Port)

		core.RegisterTCPConnHandler(socks.NewTCPHandler(proxyHost, proxyPort, fakeDns))
		core.RegisterUDPConnHandler(socks.NewUDPHandler(proxyHost, proxyPort, *args.UdpTimeout, dnsCache, fakeDns))
	})
}
