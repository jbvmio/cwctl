package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/jbvmio/cwctl"
	"github.com/spf13/cobra"
)

var filePath, dstName string
var cmdLocalFile = &cobra.Command{
	Use:   "file",
	Short: "send a local or http(s) file to Computer",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var data []byte
		var err error
		if len(args) > 0 {
			filePath = args[0]
		}
		if filePath == "" {
			Failf("error: no file specified")
		}
		if dstName == "" {
			_, dstName = filepath.Split(filePath)
			if dstName == "" {
				Failf("error: could not auto determine destination filename from %q, try --name", filePath)
			}
		}
		switch {
		case strings.HasPrefix(filePath, `http`):
			var resp *http.Response
			resp, err = http.Get(filePath)
			if err != nil {
				Failf("error retrieving file from URL %q: %v", filePath, err)
			}
			data, err = ioutil.ReadAll(resp.Body)
			resp.Body.Close()
		default:
			data, err = ioutil.ReadFile(filePath)
		}
		if err != nil {
			Failf("error reading file from %q: %v", filePath, err)
		}

		fmt.Println(">>", len(data))

		if len(data) == 0 {
			Failf("error: no data from %q", filePath)
		}
		client := initClient(cfg)
		cpu, err := cwctl.GetComputer(client, cmdFlags.ComputerID)
		if err != nil {
			Failf("error attempting GetComputer: %v", err)
		}
		if cpu.Id != cmdFlags.ComputerID {
			Failf("error validating computerID: %q doesn't match %s", cpu.Id, cmdFlags.ComputerID)
		}
		cmdFlags.UsePowerShell = false
		if cpu.IsWindows() {
			cmdFlags.UsePowerShell = true
		}
		cmdFlags.CommandText = encodedFileCommand(data, dstName, cmdFlags.UsePowerShell)
		target, err := cpu.ExecuteCommand(client, cmdFlags)
		if err != nil {
			Failf("error attempting SendLocalFile: %v", err)
		}
		handlePrint(target, outFormat)

	},
}

func init() {
	cmdLocalFile.Flags().StringVarP(&cmdFlags.ComputerID, "computer-id", `C`, "", "ID of the Computer to target.")
	cmdLocalFile.Flags().StringVarP(&cmdFlags.Directory, "dir", `D`, "", "Working Directory for the Command.")
	cmdLocalFile.Flags().StringVar(&dstName, "name", "", "Destination Filename on the Computer.")
	cmdLocalFile.MarkFlagRequired(`computer-id`)
}
