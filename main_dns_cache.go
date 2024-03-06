//go:build dnscache
// +build dnscache

package main

import (
	"flag"

	"github.com/tahirmahm123/go-tun2socks/common/dns/cache"
)

func init() {
	args.DisableDnsCache = flag.Bool("disableDNSCache", false, "Disable DNS cache (SOCKS5 and Shadowsocks handler)")

	addPostFlagsInitFn(func() {
		if *args.DisableDnsCache {
			dnsCache = nil
		} else {
			dnsCache = cache.NewSimpleDnsCache()
		}
	})
}
