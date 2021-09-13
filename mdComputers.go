package cwctl

import (
	"encoding/json"
	"fmt"

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

// GetComputers returns a list of Clients.
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

// Stats?
/*
   "CpuScore": 8.7,
   "D3DScore": 9.9,
   "DiskScore": 5.9,
   "GraphicsScore": 6.5,
   "MemoryScore": 8.7,
*/
