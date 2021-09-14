package cwctl

import (
	"encoding/json"
	"fmt"

	"github.com/jbvmio/cwctl/connectwise"
)

const (
	pwshPrefix = `powershell.exe!!! -NonInteractive -Command `
)

type CommandHistory struct {
	Id           int
	ComputerId   int
	DateExecuted connectwise.Date
	Command      string
	Status       string
	Output       string
	Parameters   string
	User         string
	DateFinished connectwise.Date
}

type CommandExecute struct {
	ComputerId int
	Command    IdNameStr
	Parameters []string
}

type CommandExecuteResponse struct {
	Id                  int
	ComputerId          int
	Command             IdNameStr
	Status              string
	Parameters          []string
	Output              string
	Fastalk             bool
	DateLastInventoried connectwise.Date
}

// GetCommandsExecuted returns Command History for a Computer.
func GetCommandsExecuted(C *connectwise.Client, params *connectwise.Parameters, computerID string) ([]CommandHistory, error) {
	var (
		resource []CommandHistory
		desc     string         = `commands executed`
		ep       connectwise.EP = connectwise.EPComputerCmdHistory
	)
	b, err := C.Get(ep, params, computerID)
	if err != nil {
		return resource, fmt.Errorf("error retrieving %s: %w", desc, err)
	}
	err = json.Unmarshal(b, &resource)
	if err != nil {
		return resource, fmt.Errorf("error unmarshaling %s: %v", desc, err)
	}
	return resource, nil
}

func ExecuteCommand(C *connectwise.Client, params *connectwise.Parameters, body CommandExecute) (CommandExecuteResponse, error) {
	var (
		resource CommandExecuteResponse
		desc     string         = `executing command`
		ep       connectwise.EP = connectwise.EPComputerCmdExec
	)
	b, err := C.Post(ep, params, body, body.ComputerId)
	if err != nil {
		return resource, fmt.Errorf("error %s: %w", desc, err)
	}
	err = json.Unmarshal(b, &resource)
	if err != nil {
		return resource, fmt.Errorf("error unmarshaling %s response: %v", desc, err)
	}
	return resource, nil
}

func WrapPwshCommand(cmd string) string {
	command := pwshPrefix
	command += `"`
	command += cmd
	command += `"`
	return command
}
