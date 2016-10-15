package main

import (
	"flag"

	"github.com/go-chat/lib"
)

func main() {
	var isHost bool
	var ipAddr string

	flag.BoolVar(&isHost, "host", false, "make this the chat host, default is guest")
	flag.StringVar(&ipAddr, "ip", "", "the ip address the host is listening on or the ip address guest is connecting to")
	flag.Parse()

	if isHost {
		lib.RunHost(ipAddr)
	} else {
		lib.RunGuest(ipAddr)
	}
}
