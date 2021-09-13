package cmd

import (
	"encoding/json"

	"github.com/jbvmio/ewctl/connectwise"
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

func getClients(C *connectwise.Client, params *connectwise.Parameters) []Client {
	b, err := C.GetClients(params)
	if err != nil {
		Failf("error retrieving clients: %v", err)
	}
	var clients []Client
	err = json.Unmarshal(b, &clients)
	if err != nil {
		Failf("error unmarshaling clients: %v", err)
	}
	return clients
}
