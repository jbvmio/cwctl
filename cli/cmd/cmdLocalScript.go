package cmd

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/jbvmio/cwctl"
	"github.com/spf13/cobra"
)

var scriptPath string
var cmdLocalScript = &cobra.Command{
	Use:     "local-script",
	Aliases: []string{`ls`, `lscript`},
	Short:   "run local script",
	Run: func(cmd *cobra.Command, args []string) {
		cmdFlags.UsePowerShell = true
		data, err := ioutil.ReadFile(scriptPath)
		if err != nil {
			Failf("error reading script: %v", err)
		}
		enc := base64.StdEncoding.EncodeToString(data)
		scriptCmd := fmt.Sprintf("iex $([System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String(\"%s\")))", enc)
		client := initClient(cfg)
		cmdFlags.CommandText = scriptCmd
		cpu, err := cwctl.GetComputer(client, cmdFlags.ComputerID)
		if err != nil {
			Failf("error attempting GetComputer: %v", err)
		}
		if cpu.Id != cmdFlags.ComputerID {
			Failf("error validating computerID: %q doesn't match %s", cpu.Id, cmdFlags.ComputerID)
		}
		if !cpu.IsWindows() {
			Failf("error: Windows Only ATM")
		}
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
	cmdLocalScript.Flags().StringVar(&scriptPath, "script", "", "Path to Script.")
	cmdLocalScript.MarkFlagRequired(`computer-id`)
	cmdLocalScript.MarkFlagRequired(`script`)
}
