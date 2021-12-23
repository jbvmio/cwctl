package cwctl

import (
	"encoding/json"
	"fmt"

	"github.com/jbvmio/cwctl/connectwise"
)

const windowsCdDirFmt = `cd %s;if(!$?){return};%s`

type CommandPrompt struct {
	ComputerID    string `json:"ComputerID,omitempty"`
	CommandID     int    `json:"CommandID,omitempty"`
	CommandText   string
	Directory     string
	RunAsAdmin    bool
	UsePowerShell bool
}

func ExecuteCommandPrompt(C *connectwise.Client, params *connectwise.Parameters, body CommandPrompt) (CommandPrompt, error) {
	var (
		resource int
		desc     string         = `executing command prompt`
		ep       connectwise.EP = connectwise.EPComputerCmdPrompt
	)
	if body.ComputerID == "" {
		return body, fmt.Errorf("error %s: %s", desc, `missing computerID`)
	}
	ID := body.ComputerID
	body.ComputerID = ""
	b, err := C.Post(ep, params, body, ID)
	body.ComputerID = ID
	if err != nil {
		return body, fmt.Errorf("error %s: %w", desc, err)
	}
	err = json.Unmarshal(b, &resource)
	if err != nil {
		return body, fmt.Errorf("error unmarshaling %s response: %v", desc, err)
	}
	body.CommandID = resource
	return body, nil
}
