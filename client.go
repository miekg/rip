package rip

import "net"

// Send sends p on the connection c.
func (p *Packet) Send(c net.Conn) error {
	buf, err := p.Pack()
	if err != nil {
		return err
	}
	n, err := c.Write(buf)
	if err != nil {
		return err
	}
	n = n
	// n != len(buf)
	return nil
}
