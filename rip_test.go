package rip

import (
	"fmt"
	"net"
	"testing"
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

func newTestPacket1() *Packet {
	p := new(Packet)
	p.Command = Response
	p.Version = One

	r1 := &Route1{Addr: net.ParseIP("127.0.0.1"), Family: 2, Metric: 1}
	p.Routes = append(p.Routes, &Route{Route1: r1})
	return p
}

func newTestPacket2() *Packet {
	p := new(Packet)
	p.Command = Response
	p.Version = Two

	r2 := &Route2{Addr: net.ParseIP("127.0.0.1"), Family: 2, Metric: 1, NextHop: net.ParseIP("127.0.0.53")}
	p.Routes = append(p.Routes, &Route{Route2: r2})
	return p
}

func TestSend(t *testing.T) {
	c, err := net.Dial("udp", fmt.Sprintf("%s:%d", "127.0.0.1", Port))
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	p1 := newTestPacket1()
	if err := p1.Send(c); err != nil {
		t.Fatal(err)
	}
	p2 := newTestPacket2()
	if err := p2.Send(c); err != nil {
		t.Fatal(err)
	}
}
