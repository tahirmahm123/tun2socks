//go:build echo
// +build echo

package main

import (
	"github.com/tahirmahm123/go-tun2socks/core"
	"github.com/tahirmahm123/go-tun2socks/proxy/echo"
)

func init() {
	registerHandlerCreater("echo", func() {
		core.RegisterTCPConnHandler(echo.NewTCPHandler())
		core.RegisterUDPConnHandler(echo.NewUDPHandler())
	})
}
