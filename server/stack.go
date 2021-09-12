package server

import (
	"WSServer/config"
	guuid "github.com/google/uuid"
	"github.com/subchen/go-log"
)

var hubs stack

type stack struct {
	items []*Hub
}

// InitStack - Initializes a stack of hubs
func InitStack() {
	defer log.Debugln("created stack of hubs")

	hubs = stack{
		items: make([]*Hub, 0),
	}
}

// NewHub - Returns a hub instance
func (s *stack) NewHub() *Hub {
	defer log.Debugln("created new hub")

	var hub Hub

	hub.ID = guuid.New()
	hub.Clients = make(map[string]*Client)

	s.items = append(s.items, &hub)

	return &hub
}

// GetActualHub - Returns a hub to which it is possible to add an entry
func (s *stack) GetActualHub() *Hub {
	sil := len(s.items)
	if sil != 0 {
		hub := s.items[sil-1]

		if len(hub.Clients) < config.Get().ClientNumber {
			return hub
		}
	}

	return s.NewHub()
}

// GetClientByID - return client by identifier
func (s *stack) GetClientByID(ID string) *Client {
	for _, item := range s.items {
		for _, client := range item.Clients {
			if client.ID.String() == ID {
				return client
			}
		}
	}

	return nil
}
