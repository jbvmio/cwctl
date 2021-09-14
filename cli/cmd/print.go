package cmd

import (
	"encoding/json"
	"fmt"

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
	}
	tbl.Print()
}

func handlePrint(object interface{}, format string) {
	switch format {
	case `raw`:
		if o, ok := object.([]byte); ok {
			fmt.Printf("%s", o)
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
	s := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		s = str[0:num] + "..."
	}
	return s
}
