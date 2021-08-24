package wscontroller

import (
	"context"
)

// TODO: close message
var closeMessage = []byte("close")

// Hub maintains the set of active sessions and broadcasts messages to the
// sessions.
type Hub struct {
	// Registered sessions.
	sessions map[*Session]struct{}

	// Inbound messages from the sessions.
	broadcastCh chan []byte

	// Register requests from the sessions.
	registerCh chan *Session

	// Unregister requests from sessions.
	unregisterCh chan *Session
}

func NewHub() *Hub {
	return &Hub{
		broadcastCh:  make(chan []byte),
		registerCh:   make(chan *Session),
		unregisterCh: make(chan *Session),
		sessions:     make(map[*Session]struct{}),
	}
}

func (h *Hub) Broadcast(message []byte) {
	h.broadcastCh <- message
}

func (h *Hub) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			h.broadcast(closeMessage)
			return ctx.Err()
		case session := <-h.registerCh:
			h.sessions[session] = struct{}{}
		case session := <-h.unregisterCh:
			if _, ok := h.sessions[session]; ok {
				delete(h.sessions, session)
				close(session.send)
			}
		case message := <-h.broadcastCh:
			h.broadcast(message)
		}
	}
}

func (h *Hub) broadcast(message []byte) {
	for session := range h.sessions {
		select {
		case session.send <- message:
		default:
			close(session.send)
			delete(h.sessions, session)
		}
	}
}
