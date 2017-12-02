package rip

import (
	"encoding/binary"
	"net"
)

// unpackHeader unpacks an header, returning the offset to the end of the header.
func unpackHeader(packet []byte, off int) (h Header, off1 int, err error) {
	if off == len(packet) {
		return h, off, nil
	}

	h.Command, off, err = unpackUint8(packet, off)
	if err != nil {
		return h, len(packet), err
	}
	h.Version, off, err = unpackUint8(packet, off)
	if err != nil {
		return h, len(packet), err
	}
	h.mbz, off, err = unpackUint16(packet, off)
	if err != nil {
		return h, len(packet), err
	}
	return h, off, nil
}

// pack packs an header, returning the offset to the end of the header.
func (h Header) pack(packet []byte, off int) (off1 int, err error) {
	if off == len(packet) {
		return off, nil
	}

	off, err = packUint8(h.Command, packet, off)
	if err != nil {
		return len(packet), err
	}
	off, err = packUint8(h.Version, packet, off)
	if err != nil {
		return len(packet), err
	}
	off, err = packUint16(h.mbz, packet, off)
	if err != nil {
		return len(packet), err
	}
	return off, nil
}

func unpackRoute1(packet []byte, off int) (r *Route1, off1 int, err error) {
	if off == len(packet) {
		return r, off, nil
	}
	r = new(Route1)

	r.Family, off, err = unpackUint16(packet, off)
	if err != nil {
		return r, len(packet), err
	}
	r.mbz1, off, err = unpackUint16(packet, off)
	if err != nil {
		return r, len(packet), err
	}
	r.Addr, off, err = unpackIP(packet, off)
	if err != nil {
		return r, len(packet), err
	}
	r.mbz2, off, err = unpackUint32(packet, off)
	if err != nil {
		return r, len(packet), err
	}
	r.mbz3, off, err = unpackUint32(packet, off)
	if err != nil {
		return r, len(packet), err
	}
	r.Metric, off, err = unpackUint32(packet, off)
	if err != nil {
		return r, len(packet), err
	}
	return r, off, nil
}

func (r *Route2) pack(packet []byte, off int) (off1 int, err error) {
	if off == len(packet) {
		return off, nil
	}
	r = new(Route2)

	off, err = packUint16(r.Family, packet, off)
	if err != nil {
		return len(packet), err
	}
	off, err = packUint16(r.RouteTag, packet, off)
	if err != nil {
		return len(packet), err
	}
	off, err = packIP(r.Addr, packet, off)
	if err != nil {
		return len(packet), err
	}
	off, err = packUint32(r.Mask, packet, off)
	if err != nil {
		return len(packet), err
	}
	off, err = packIP(r.NextHop, packet, off)
	if err != nil {
		return len(packet), err
	}
	off, err = packUint32(r.Metric, packet, off)
	if err != nil {
		return len(packet), err
	}
	return off, nil
}

func unpackRoute2(packet []byte, off int) (r *Route2, off1 int, err error) {
	if off == len(packet) {
		return r, off, nil
	}
	r = new(Route2)

	r.Family, off, err = unpackUint16(packet, off)
	if err != nil {
		return r, len(packet), err
	}
	r.RouteTag, off, err = unpackUint16(packet, off)
	if err != nil {
		return r, len(packet), err
	}
	r.Addr, off, err = unpackIP(packet, off)
	if err != nil {
		return r, len(packet), err
	}
	r.Mask, off, err = unpackUint32(packet, off)
	if err != nil {
		return r, len(packet), err
	}
	r.NextHop, off, err = unpackIP(packet, off)
	if err != nil {
		return r, len(packet), err
	}
	r.Metric, off, err = unpackUint32(packet, off)
	if err != nil {
		return r, len(packet), err
	}
	return r, off, nil
}

func (r *Route1) pack(packet []byte, off int) (off1 int, err error) {
	if off == len(packet) {
		return off, nil
	}
	off, err = packUint16(r.Family, packet, off)
	if err != nil {
		return len(packet), err
	}
	off, err = packUint16(r.mbz1, packet, off)
	if err != nil {
		return len(packet), err
	}
	off, err = packIP(r.Addr, packet, off)
	if err != nil {
		return len(packet), err
	}
	off, err = packUint32(r.mbz2, packet, off)
	if err != nil {
		return len(packet), err
	}
	off, err = packUint32(r.mbz3, packet, off)
	if err != nil {
		return len(packet), err
	}
	off, err = packUint32(r.Metric, packet, off)
	if err != nil {
		return len(packet), err
	}
	return off, nil

}

func unpackUint8(packet []byte, off int) (i uint8, off1 int, err error) {
	if off+1 > len(packet) {
		return 0, len(packet), &PackError{err: "overflow unpacking uint8"}
	}
	return uint8(packet[off]), off + 1, nil
}

func packUint8(i uint8, packet []byte, off int) (off1 int, err error) {
	if off+1 > len(packet) {
		return len(packet), &PackError{err: "overflow packing uint8"}
	}
	packet[off] = byte(i)
	return off + 1, nil
}

func unpackUint16(packet []byte, off int) (i uint16, off1 int, err error) {
	if off+2 > len(packet) {
		return 0, len(packet), &PackError{err: "overflow unpacking uint16"}
	}
	return binary.BigEndian.Uint16(packet[off:]), off + 2, nil
}

func packUint16(i uint16, packet []byte, off int) (off1 int, err error) {
	if off+2 > len(packet) {
		return len(packet), &PackError{err: "overflow packing uint16"}
	}
	binary.BigEndian.PutUint16(packet[off:], i)
	return off + 2, nil
}

func unpackUint32(packet []byte, off int) (i uint32, off1 int, err error) {
	if off+4 > len(packet) {
		return 0, len(packet), &PackError{err: "overflow unpacking uint32"}
	}
	return binary.BigEndian.Uint32(packet[off:]), off + 4, nil
}

func packUint32(i uint32, packet []byte, off int) (off1 int, err error) {
	if off+4 > len(packet) {
		return len(packet), &PackError{err: "overflow packing uint32"}
	}
	binary.BigEndian.PutUint32(packet[off:], i)
	return off + 4, nil
}

func unpackIP(msg []byte, off int) (net.IP, int, error) {
	if off+net.IPv4len > len(msg) {
		return nil, len(msg), &PackError{err: "overflow unpacking Addr"}
	}
	a := append(make(net.IP, 0, net.IPv4len), msg[off:off+net.IPv4len]...)
	off += net.IPv4len
	return a, off, nil
}

func packIP(a net.IP, msg []byte, off int) (int, error) {
	// It must be a slice of 4, even if it is 16, we encode only the first 4
	if off+net.IPv4len > len(msg) {
		return len(msg), &PackError{err: "overflow packing Addr"}
	}

	switch len(a) {
	case net.IPv4len, net.IPv6len:
		copy(msg[off:], a.To4())
		off += net.IPv4len
	case 0:
		// Not set, skip it
		off += net.IPv4len
	default:
		return len(msg), &PackError{err: "overflow packing Addr"}
	}
	return off, nil
}
