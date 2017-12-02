package rip

import (
	"net"
	"testing"
)

func TestPackRIP1(t *testing.T) {
	p := new(Packet)
	p.Command = Response
	p.Version = One

	p.Routes = append(p.Routes, &Route{Route1: &Route1{Addr: net.ParseIP("127.0.0.1")}})

	buf, err := p.Pack()
	if err != nil {
		t.Fatal(err)
	}
	if len(buf) != p.Len() {
		t.Errorf("buf and packet len don't match: %d != %d", len(buf), p.Len())
	}

	p1, err := Unpack(buf)
	if err != nil {
		t.Fatal(err)
	}
	if p1.String() != p.String() {
		t.Errorf("Pack/Unpack lost information, strings don't match")
		t.Logf("\noriginal\n%s\nnew\n%s", p, p1)
	}
}
