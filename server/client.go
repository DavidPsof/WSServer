package server

import guuid "github.com/google/uuid"

// Client - The structure describing the client
type Client struct {
	ID           guuid.UUID
	HubID        string
	ConnectionID string
}

// NewClient - return example of client
func NewClient(CID string, hID string) Client {
	return Client{
		ID:           guuid.New(),
		HubID:        hID,
		ConnectionID: CID,
	}
}
