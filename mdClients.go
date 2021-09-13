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
	Locations  []IdNameInt
}

// GetClients returns a list of Clients.
func GetClients(C *connectwise.Client, params *connectwise.Parameters) ([]Client, error) {
	var (
		resource []Client
		desc     string         = `clients`
		ep       connectwise.EP = connectwise.EPClients
	)
	b, err := C.Get(ep, params)
	if err != nil {
		return resource, fmt.Errorf("error retrieving %s: %w", desc, err)
	}
	err = json.Unmarshal(b, &resource)
	if err != nil {
		return resource, fmt.Errorf("error unmarshaling %s: %v", desc, err)
	}
	return resource, nil
}
