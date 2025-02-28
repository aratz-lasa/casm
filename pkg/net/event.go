package net

import (
	"fmt"

	"github.com/libp2p/go-libp2p-core/peer"
)

const (
	// EventJoined indicates a peer has joined the neighborhood.
	EventJoined Event = iota

	// EventLeft indicates a peer has left the neighborhood.
	EventLeft
)

// EvtState is an event that signals the state of the overlay network.
type EvtState struct {
	Event Event
	Peer  peer.ID
	es    edgeMap
}

func (ev EvtState) String() string { return fmt.Sprintf("%s %s", ev.Peer, ev.Event) }

func (ev EvtState) Loggable() map[string]interface{} {
	return map[string]interface{}{
		"peer":  ev.Peer,
		"event": ev.Event,
		"edges": ev.Edges(),
	}
}

func (ev EvtState) Edges() peer.IDSlice { return ev.es.Slice() }

// An Event represents a state transition in a neighborhood.
type Event uint8

func (e Event) Loggable() map[string]interface{} {
	return map[string]interface{}{
		"event": e.String(),
	}
}

func (e Event) String() string {
	switch e {
	case EventJoined:
		return "joined"
	case EventLeft:
		return "left"
	}

	panic(fmt.Sprintf("invalid event '%d'", e))
}
