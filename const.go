package rip

// Maximum number of entries in a packet. Minimum is 1.
const maxRoutes = 25

// Various contants for use in RIP
const (
	One  = 1      // One is the version used for RIP-1.
	Two  = 2      // Two is the version used for RIP-2.
	Auth = 0xFFFF // If the address family contains this value the whole packet should be used for auth. See RFC 2453, Section 4.1.

	Request  = 1 // Header commands with this value are requests.
	Response = 2 // Header commands with this value are Responses.

	Infinity = 16 // Infinity in the Metrics.
)
