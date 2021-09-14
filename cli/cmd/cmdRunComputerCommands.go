package cmd

import (
	"github.com/jbvmio/cwctl"
	"github.com/spf13/cobra"
)

var cliIntID int
var cmdRunComputerCommands = &cobra.Command{
	Use:     "commands",
	Aliases: []string{"command", "command-execute"},
	Args:    cobra.MinimumNArgs(1),
	Short:   "run / execute computer command",
	Run: func(cmd *cobra.Command, args []string) {
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
		target, err := cwctl.ExecuteCommand(client, paramsDefault.merge(&cliFlags), body)
		if err != nil {
			Failf("error attempting RunComputerCommands: %v", err)
		}
		handlePrint(target, outFormat)
	},
}

func init() {
	cmdRunComputerCommands.Flags().IntVarP(&cliIntID, "computer-id", `C`, cliIntID, "ID of the Computer to target.")
	cmdRunComputerCommands.MarkFlagRequired(`computer-id`)
}
