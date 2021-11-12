package cwctl

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/jbvmio/cwctl/connectwise"
)

type Computer struct {
	Id                       string
	ComputerName             string
	LocalIPAddress           string
	GatewayIPAddress         string
	SystemUptime             int64
	MACAddress               string
	SerialNumber             string
	DateAdded                connectwise.Date
	UserIdleTime             int64
	LastStartup              connectwise.Date
	OperatingSystemName      string
	OperatingSystemVersion   string
	DomainName               string
	DomainNameServers        []string
	RemoteAgentVersion       string
	RemoteAgentLastContact   connectwise.Date
	RemoteAgentLastInventory connectwise.Date
	LastInventoryReceived    connectwise.Date
	WindowsUpdateDate        connectwise.Date
	AntivirusDefinitionDate  connectwise.Date
	LastHeartbeat            connectwise.Date
	LastUserName             string
	Bandwidth                int64
	BandwidthDisplay         string
	AssetDate                connectwise.Date
	AssetTag                 string
	Type                     string
	Status                   string
	MasterMode               string
	VirusScanner             IdNameStr
	IsFasTalk                bool
	IsMaster                 bool
	IsNetworkProbe           bool
	IsHeartbeatEnabled       bool
	IsMaintenanceModeEnabled bool
	IsTunnelSupported        bool
	IsVirtualMachine         bool
	IsLockedDown             bool
	IsSystemAccount          bool
	IsRebootNeeded           bool
	IsVirtualHost            bool
	IsHeartbeatRunning       bool
	HasIntelVPRO             bool
	HasIntelAMT              bool
	HasHPiLO                 bool
	TempFiles                string
	Comment                  string
	OpenPortsTCP             []int
	OpenPortsUDP             []int
	TotalMemory              int64
	FreeMemory               int64
	LoggedInUsers            []LoggedInUser
	UserAccounts             []string
	Location                 IdNameInt
	Client                   Client
}

type LoggedInUser struct {
	LoggedInUserName string
	ConsoleId        int
}

func (c *Computer) IsWindows() bool {
	switch {
	case strings.Contains(c.OperatingSystemName, `Windows`):
		return true
	case strings.Contains(c.OperatingSystemName, `Microsoft`):
		return true
	}
	return false
}

func (c *Computer) IsOnline() bool {
	return c.Status == `Online`
}

func (c *Computer) ExecuteCommand(C *connectwise.Client, cmd CommandPrompt) (CommandPrompt, error) {
	cmd.ComputerID = c.Id
	switch {
	case c.IsWindows():
		if cmd.Directory == "" {
			cmd.Directory = `%windir%\\system32`
		}
	default:
		if cmd.Directory == "" {
			cmd.Directory = `/tmp/`
		}
		cmd.UsePowerShell = false
		cmd.RunAsAdmin = false
	}
	return ExecuteCommandPrompt(C, nil, cmd)
}

// GetComputers returns a list of Computers.
func GetComputers(C *connectwise.Client, params *connectwise.Parameters) ([]Computer, error) {
	var (
		resource []Computer
		desc     string         = `computers`
		ep       connectwise.EP = connectwise.EPComputers
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

// GetComputer returns a single Computer.
func GetComputer(C *connectwise.Client, computerID string) (Computer, error) {
	var (
		resource Computer
		desc     string         = `computer`
		ep       connectwise.EP = connectwise.EPComputer
	)
	if !ValidComputerID(computerID) {
		return resource, fmt.Errorf("error retrieving %s %q: %s", desc, computerID, `invalid computerID`)
	}
	b, err := C.Get(ep, nil, computerID)
	if err != nil {
		return resource, fmt.Errorf("error retrieving %s %q: %w", desc, computerID, err)
	}
	err = json.Unmarshal(b, &resource)
	if err != nil {
		return resource, fmt.Errorf("error unmarshaling %s %q: %w", desc, computerID, err)
	}
	return resource, nil
}

func ValidComputerID(computerID string) bool {
	valid, _ := regexp.MatchString(`^[0-9]+$`, computerID)
	return valid
}
