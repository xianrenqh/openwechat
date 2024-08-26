//go:build !(go1.20 || go1.21 || go1.22)

package main

import "time"

type entry struct {
	Name       string
	Value      string
	Quoted     bool
	Domain     string
	Path       string
	SameSite   string
	Secure     bool
	HttpOnly   bool
	Persistent bool
	HostOnly   bool
	Expires    time.Time
	Creation   time.Time
	LastAccess time.Time

	// seqNum is a sequence number so that Cookies returns cookies in a
	// deterministic order, even for cookies that have equal Path length and
	// equal Creation time. This simplifies testing.
	seqNum uint64 // nolint:unused
}
