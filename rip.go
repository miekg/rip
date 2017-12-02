package rip

import (
	"fmt"
	"net"
)

// Packet is a packet that is being exchanged in the RIP protocol. This contains
// either version 1 or version 2 Route Entries.
type Packet struct {
	Header

	// Routes contains all the routes. Each *Route either contains a *Route1
	// or a Route2 depending in the version in the header.
	Routes []*Route
}

// Len returns the length of p in octets.
func (p *Packet) Len() int {
	l := 4
	for _, r := range p.Routes {
		if r.Route1 != nil {
			l += r.Route1.len()
			continue
		}
		if r.Route2 != nil {
			l += r.Route2.len()
		}
	}
	return l
}

// Header is the packet's Header in the RIP version 2 protocol.
type Header struct {
	Command uint8
	Version uint8
	mbz     uint16
}

// Route is a RIP-1 or RIP-2 Route Entry.
type Route struct {
	*Route1
	*Route2
}

// Route1 is a RIP-1 Route Entry.
type Route1 struct {
	Family uint16
	mbz1   uint16
	Addr   net.IP
	mbz2   uint32
	mbz3   uint32
	Metric uint32
}

func (r1 *Route1) len() int { return 20 }

// Route2 is a RIP-2 Route Entry.
type Route2 struct {
	Family   uint16
	RouteTag uint16
	Addr     net.IP
	Mask     uint32
	NextHop  uint32
	Metric   uint32
}

func (r2 *Route2) len() int { return 20 }

// Authentication is used for authentication purposes. This has not been implemented.
type Authentication struct {
	Password string
}

// Unpack unpacks packet into a Packet from the network.
func Unpack(packet []byte) (*Packet, error) {
	p := new(Packet)
	h, off, err := unpackHeader(packet, 0)
	if err != nil {
		return nil, err
	}
	p.Header = h

	routes := []*Route{}
	i := 0

Unpack:
	switch p.Header.Version {
	case One:
		for i = 0; i < maxRoutes; i++ {
			if off == len(packet) {
				break Unpack
			}
			if off > len(packet) {
				return nil, &PackError{err: "overflow unpacking packet"}
			}

			r1 := new(Route1)
			r1, off, err = unpackRoute1(packet, off)
			if err != nil {
				return nil, err
			}
			routes = append(routes, &Route{Route1: r1})
		}

	case Two:
		for i = 0; i < maxRoutes; i++ {
			if off == len(packet) {
				break Unpack
			}
			if off > len(packet) {
				return nil, &PackError{err: "overflow unpacking packet"}
			}

			r2 := new(Route2)
			r2, off, err = unpackRoute2(packet, off)
			if err != nil {
				return nil, err
			}
			routes = append(routes, &Route{Route2: r2})
		}
	default:
		return nil, &ProtoError{err: fmt.Sprintf("%s: %d", "bad rip packet: wrong version", p.Header.Version)}
	}

	if i == 0 {
		return nil, &ProtoError{err: "bad rip packet: 0 route entries"}
	}

	return p, nil
}

// Pack packs a Packet so that it can be send on the network.
func (p *Packet) Pack() (packet []byte, err error) {
	if len(p.Routes) > maxRoutes {
		return nil, &ProtoError{err: fmt.Sprintf("bad rip packet: %d route entries", len(p.Routes))}
	}
	if p.Header.Version != One && p.Header.Version != Two {
		return nil, &ProtoError{err: fmt.Sprintf("%s: %d", "bad rip packet: wrong version", p.Header.Version)}
	}

	off := 0
	packet = make([]byte, p.Len())

	off, err = p.Header.pack(packet, 0)
	switch p.Header.Version {
	case One:
		for _, r := range p.Routes {
			off, err = r.Route1.pack(packet, off)
			if err != nil {
				return nil, err
			}
		}
	case Two:
		for _, r := range p.Routes {
			off, err = r.Route2.pack(packet, off)
			if err != nil {
				return nil, err
			}
		}
	}

	return packet, nil
}
