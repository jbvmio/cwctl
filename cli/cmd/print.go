package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jbvmio/cwctl"
	"github.com/jbvmio/cwctl/connectwise"
	"github.com/rodaine/table"
	"github.com/tidwall/pretty"
	"gopkg.in/yaml.v3"
)

func printOut(i interface{}) {
	var tbl table.Table
	switch i := i.(type) {
	case string:
		tbl = table.New("OBJECT")
		tbl.AddRow(i)
	case []connectwise.EndPoint:
		tbl = table.New("ID", "PATH")
		for _, v := range i {
			tbl.AddRow(v.ID, v.Path)
		}
	case []cwctl.Client:
		tbl = table.New("ID", "NAME", "COMPANY", "CITY", "STATE", "LOCATIONS")
		for _, v := range i {
			tbl.AddRow(v.Id, v.Name, v.Company, v.City, v.State, len(v.Locations))
		}
	case []cwctl.Computer:
		tbl = table.New("ID", "DOMAIN", "COMPUTER", "IP", "OS", "OSVersion", "VirusScanner", "AntivirusDefinitions", "AGENT", "LastHeartBeat", "LastUser")
		for _, v := range i {
			tbl.AddRow(v.Id, v.DomainName, v.ComputerName, v.LocalIPAddress, v.OperatingSystemName, v.OperatingSystemVersion, v.VirusScanner.Name, v.AntivirusDefinitionDate, v.RemoteAgentVersion, v.LastHeartbeat, v.LastUserName)
		}
	case []cwctl.Command:
		tbl = table.New("ID", "NAME", "DESCRIPTION")
		for _, v := range i {
			tbl.AddRow(v.Id, v.Name, truncateString(v.Description, 120))
		}
	case []cwctl.CommandHistory:
		tbl = table.New("ID", "ComputerID", "STATUS", "USER", "PARAMETERS", "EXECUTED", "FINISHED")
		for _, v := range i {
			tbl.AddRow(v.Id, v.ComputerId, v.Status, v.User, truncateString(v.Parameters, 120), v.DateExecuted, v.DateFinished)
		}
	case cwctl.CommandExecuteResponse:
		tbl = table.New("ID", "ComputerID", "STATUS", "PARAMETERS", "LastInventoried")
		tbl.AddRow(i.Id, i.ComputerId, i.Status, len(i.Parameters), i.DateLastInventoried)
	case []cwctl.CommandExecuteResponse:
		tbl = table.New("ID", "ComputerID", "STATUS", "PARAMETERS", "LastInventoried")
		for _, v := range i {
			tbl.AddRow(v.Id, v.ComputerId, v.Status, len(v.Parameters), v.DateLastInventoried)
		}
	}
	tbl.Print()
}

func handlePrint(object interface{}, format string) {
	switch format {
	case `raw`:
		if o, ok := object.([]byte); ok {
			fmt.Printf("%s\n", o)
			return
		}
		Failf("unable to display, not raw object")
	case `yaml`:
		fmtString, err := yaml.Marshal(object)
		if err != nil {
			Failf("unable to format yaml: %v", err)
		}
		fmt.Printf("%s", fmtString)
	case `json`:
		fmtString, err := json.Marshal(object)
		if err != nil {
			Failf("unable to format json: %v", err)
		}
		fmt.Printf("%s", fmtString)
	case `pretty`:
		fmtString, err := json.Marshal(object)
		if err != nil {
			Failf("unable to format json: %v", err)
		}
		fmt.Printf("%s", pretty.Pretty(fmtString))
	default:
		printOut(object)
	}
}

func truncateString(str string, num int) string {
	s := strings.ReplaceAll(str, `\r\n`, `\n`)
	s = strings.ReplaceAll(s, `\n`, ``)
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		s = str[0:num] + "..."
	}
	return s
}
