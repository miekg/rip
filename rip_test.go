package rip

import (
	"fmt"
	"net"
)

func ExampleStringer() {
	p := new(Packet)
	p.Command = Response
	p.Version = One

	p.Routes = append(p.Routes, &Route{Route1: &Route1{Addr: net.ParseIP("127.0.0.1")}})

	fmt.Println(p.String())
	// Output:
	// ;; ->>HEADER<<- command: response, version: 1
	// 1. 127.0.0.1
}
