package rip

import "net"

// Packet is a packet that is being exchanged in the RIP version 2 protocol.
type Packet struct {
	Header
	Entries []Entry
}

// Header is the packet's Header in the RIP version 2 protocol.
type Header struct {
	Command uint8
	Version uint8
	mbz     uint16
}

type Entry struct {
	Family uint16
	mbz1   uint16
	Addr   net.IP
	mbz2   uint32
	mbz3   uint32
	Metric uint32
}

const maxEntries = 25

// PackError is an error that is returned when packing or unpacking a packet.
type PackError struct {
	err string
}

func (p *PackError) Error() string { return p.err }
