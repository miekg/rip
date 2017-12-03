package rip

import (
	"fmt"
	"net"
	"testing"
)

func ExampleStringer2() {
	p := new(Packet)
	p.Command = Response
	p.Version = Two

	r2 := &Route2{Addr: net.ParseIP("127.0.0.1"), Family: 2, Metric: 1, NextHop: net.ParseIP("127.0.0.53")}
	p.Routes = append(p.Routes, &Route{Route2: r2})

	fmt.Println(p.String())
	// Output:
	// ;; ->>HEADER<<- command: response, version: 2
	// 1. 127.0.0.1 ->127.0.0.53
}

func packet2() *Packet {
	p := new(Packet)
	p.Command = Response
	p.Version = Two

	r2 := &Route2{Addr: net.ParseIP("127.0.0.1"), Family: 2, Metric: 1, NextHop: net.ParseIP("127.0.0.53")}
	p.Routes = append(p.Routes, &Route{Route2: r2})
	return p
}

func TestSend2(t *testing.T) {
	c, err := net.Dial("udp", fmt.Sprintf("%s:%d", "127.0.0.1", Port))
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	p2 := packet2()
	if err := p2.Send(c); err != nil {
		t.Fatal(err)
	}
}
