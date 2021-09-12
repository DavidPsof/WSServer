package server

import (
	guuid "github.com/google/uuid"
)

// Hub - The structure describing the hub
type Hub struct {
	ID      guuid.UUID
	Clients map[string]*Client
}

// AddClient - adds a client to the hub
func (h *Hub) AddClient(c *Client) {
	h.Clients[c.ID.String()] = c
}
