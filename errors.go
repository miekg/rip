package rip

// PackError is an error that is returned when packing or unpacking a packet.
type PackError struct {
	err string
}

func (p *PackError) Error() string { return p.err }

// ProtoError is an error that is returned when the semantics of the packet are wrong.
type ProtoError struct {
	err string
}

func (p *ProtoError) Error() string { return p.err }
