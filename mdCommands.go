package cwctl

import (
	"encoding/json"
	"fmt"

	"github.com/jbvmio/cwctl/connectwise"
)

type Command struct {
	Id          string
	Name        string
	Description string
	Level       int
}

// GetCommands returns a list of Clients.
func GetCommands(C *connectwise.Client, params *connectwise.Parameters) ([]Command, error) {
	var (
		resource []Command
		desc     string         = `commands`
		ep       connectwise.EP = connectwise.EPCommands
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
