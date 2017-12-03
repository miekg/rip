package rip

import (
	"fmt"
	"net"
	"testing"
)

func ExampleStringer1() {
	p := new(Packet)
	p.Command = Response
	p.Version = One

	p.Routes = append(p.Routes, NewRoute(net.ParseIP("127.0.0.1"), nil, 0, 0, 0))

	fmt.Println(p.String())
	// Output:
	// ;; ->>HEADER<<- command: response, version: 1
	// 1. 127.0.0.1
}

func packet1() *Packet {
	p := New(Response, One)
	p.Routes = append(p.Routes, NewRoute(net.ParseIP("127.0.0.1"), nil, 0, 1, 0))
	return p
}

func TestSend1(t *testing.T) {
	c, err := net.Dial("udp", fmt.Sprintf("%s:%d", "127.0.0.1", Port))
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	p1 := packet1()
	if err := p1.Send(c); err != nil {
		t.Fatal(err)
	}
}
