package cmd

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/jbvmio/cwctl"
	"github.com/spf13/cobra"
)

var scriptPath string
var cmdLocalScript = &cobra.Command{
	Use:   "script",
	Short: "run local or http(s) script",
	Run: func(cmd *cobra.Command, args []string) {
		var data []byte
		var err error
		switch {
		case strings.HasPrefix(scriptPath, `http`):
			var resp *http.Response
			resp, err = http.Get(scriptPath)
			if err != nil {
				Failf("error retrieving script from URL %q: %v", scriptPath, err)
			}
			data, err = ioutil.ReadAll(resp.Body)
			resp.Body.Close()
		default:
			data, err = ioutil.ReadFile(scriptPath)
		}
		if err != nil {
			Failf("error reading script from %q: %v", scriptPath, err)
		}
		switch {
		case len(data) == 0:
			Failf("error: no data from %q", filePath)
		case len(data) > maxSize:
			Failf("error: filesize of %d for %q exceeds 23k", len(data)/1024, filePath)
		}
		client := initClient(cfg)
		cpu, err := cwctl.GetComputer(client, cmdFlags.ComputerID)
		if err != nil {
			Failf("error attempting GetComputer: %v", err)
		}
		if cpu.Id != cmdFlags.ComputerID {
			Failf("error validating computerID: %qk doesn't match %s", cpu.Id, cmdFlags.ComputerID)
		}
		cmdFlags.UsePowerShell = false
		if cpu.IsWindows() {
			cmdFlags.UsePowerShell = true
		}
		cmdFlags.CommandText = encodedScriptCommand(data, cmdFlags.UsePowerShell)
		target, err := cpu.ExecuteCommand(client, cmdFlags)
		if err != nil {
			Failf("error attempting RunLocalScript: %v", err)
		}
		handlePrint(target, outFormat)
	},
}

func init() {
	cmdLocalScript.Flags().StringVarP(&cmdFlags.ComputerID, "computer-id", `C`, "", "ID of the Computer to target.")
	cmdLocalScript.Flags().StringVarP(&cmdFlags.Directory, "dir", `D`, "", "Working Directory for the Command.")
	cmdLocalScript.Flags().StringVar(&scriptPath, "script", "", "Local Path or URL to Script.")
	cmdLocalScript.MarkFlagRequired(`computer-id`)
	cmdLocalScript.MarkFlagRequired(`script`)
}
