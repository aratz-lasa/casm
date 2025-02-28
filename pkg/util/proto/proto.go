// Package protoutil contains utilities for working with libp2p's protocol.ID.
package protoutil

import (
	"strings"

	"github.com/libp2p/go-libp2p-core/protocol"
)

const sep = "/"

// Join multiple protocol.IDs into one, using the path separator.
func Join(ids ...protocol.ID) protocol.ID {
	ss := protocol.ConvertToStrings(ids)
	return protocol.ID(strings.Join(ss, sep))
}

// Parts splits the protocol.ID into its constituent components
func Parts(id protocol.ID) []protocol.ID {
	var b strings.Builder

	parts := make([]string, 0, 8) // best-effort preallocation
	for _, r := range id {
		if r == '/' {
			if b.Len() != 0 {
				parts = append(parts, b.String())
				b.Reset()
			}
			continue
		}

		b.WriteRune(r)
	}

	if b.Len() != 0 {
		parts = append(parts, b.String())
	}

	return protocol.ConvertFromStrings(parts)
}

// Split the protocol into its base path and its end-component.
func Split(id protocol.ID) (base, end protocol.ID) {
	switch parts := Parts(id); len(parts) {
	case 0:
		return
	case 1:
		return "", parts[0]
	default:
		end = parts[len(parts)-1]
		parts = parts[:len(parts)-1]
		base = "/" + Join(parts...)
		return
	}
}
