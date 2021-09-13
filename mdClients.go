package cwctl

import (
	"encoding/json"
	"fmt"

	"github.com/jbvmio/cwctl/connectwise"
)

type Client struct {
	Id         string
	ExternalId string
	Name       string
	Company    string
	City       string
	State      string
	Locations  []Location
}

type Location struct {
	Id   int
	Name string
}

// GetClients returns a list of Clients.
func GetClients(C *connectwise.Client, params *connectwise.Parameters) ([]Client, error) {
	var clients []Client
	b, err := C.GetClients(params)
	if err != nil {
		return clients, fmt.Errorf("error retrieving clients: %w", err)
	}
	err = json.Unmarshal(b, &clients)
	if err != nil {
		return clients, fmt.Errorf("error unmarshaling clients: %v", err)
	}
	return clients, nil
}
