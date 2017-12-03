package rip

import (
	"net"
)

// Route is a RIP-1 or RIP-2 Route Entry.
type Route struct {
	*Route1
	*Route2
}

// NewRoute returns a new pointer to Route. The family will be set to AF_INET (2),
// if NextHop, mask or tag is set a RIP2 route is returned, otherwise a RIP1 one.
func NewRoute(addr, nexthop net.IP, mask, metric uint32, tag uint16) *Route {
	r := new(Route)
	if nexthop != nil || mask > 0 || tag > 0 {
		r2 := &Route2{Addr: addr, NextHop: nexthop, Mask: mask, RouteTag: tag, Metric: metric, Family: 2}
		r.Route2 = r2
		return r
	}

	r1 := &Route1{Addr: addr, Metric: metric, Family: 2}
	r.Route1 = r1
	return r
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

func (r1 *Route1) String() string {
	if r1 == nil {
		return ""
	}
	s := r1.Addr.String()
	return s + "\n"
}

// Route2 is a RIP-2 Route Entry.
type Route2 struct {
	Family   uint16
	RouteTag uint16
	Addr     net.IP
	Mask     uint32
	NextHop  net.IP
	Metric   uint32
}

func (r2 *Route2) len() int { return 20 }

func (r2 *Route2) String() string {
	if r2 == nil {
		return ""
	}
	s := r2.Addr.String()
	if r2.NextHop != nil {
		s += " ->" + r2.NextHop.String()
	}
	return s + "\n"
}
