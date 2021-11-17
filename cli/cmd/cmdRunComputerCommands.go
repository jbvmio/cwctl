package cmd

import (
	"github.com/jbvmio/cwctl"
	"github.com/spf13/cobra"
)

var (
	cmdFlags cwctl.CommandPrompt
)
var cmdRunComputerCommands = &cobra.Command{
	Use:     "commands",
	Aliases: []string{"command", "command-execute"},
	Args:    cobra.MinimumNArgs(1),
	Short:   "run / execute computer command",
	Run: func(cmd *cobra.Command, args []string) {
		client := initClient(cfg)
		cmdFlags.CommandText = args[0]
		cpu, err := cwctl.GetComputer(client, cmdFlags.ComputerID)
		if err != nil {
			Failf("error attempting GetComputer: %v", err)
		}
		if cpu.Id != cmdFlags.ComputerID {
			Failf("error validating computerID: %q doesn't match %s", cpu.Id, cmdFlags.ComputerID)
		}
		target, err := cpu.ExecuteCommand(client, cmdFlags)
		if err != nil {
			Failf("error attempting RunComputerCommands: %v", err)
		}
		handlePrint(target, outFormat)
	},
}

func init() {
	//cmdRunComputerCommands.Flags().IntVarP(&cliIntID, "computer-id", `C`, cliIntID, "ID of the Computer to target.") << was used with cwctl.ExecuteCommand()
	cmdRunComputerCommands.Flags().StringVarP(&cmdFlags.ComputerID, "computer-id", `C`, "", "ID of the Computer to target.")
	cmdRunComputerCommands.Flags().StringVarP(&cmdFlags.Directory, "dir", `D`, "", "Working Directory for the Command.")
	cmdRunComputerCommands.Flags().BoolVar(&cmdFlags.UsePowerShell, "pwsh", false, "Use PowerShell.")
	cmdRunComputerCommands.MarkFlagRequired(`computer-id`)
}

// runCommandOld is ExecuteCommand, replaced by ExecuteCommandPrompt
/*
func runCommandOld(cliIntID int, args ...string) {
	commands := make([]string, len(args))
	for i := 0; i < len(args); i++ {
		commands[i] = cwctl.WrapPwshCommand(args[i])
	}
	body := cwctl.CommandExecute{
		ComputerId: cliIntID,
		Command: cwctl.IdNameStr{
			Id: `2`,
		},
		Parameters: commands,
	}
	client := initClient(cfg)
	target, err := cwctl.ExecuteCommand(client, paramsDefault.merge(&paramFlags), body)
	if err != nil {
		Failf("error attempting RunComputerCommands: %v", err)
	}
	handlePrint(target, outFormat)
}
*/
