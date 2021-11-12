package cmd

import (
	"github.com/jbvmio/cwctl"
	"github.com/spf13/cobra"
)

var cmdComputerCommands = &cobra.Command{
	Use:     "commands",
	Aliases: []string{"command", "command-history"},
	Short:   "get computer command history / details",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			cliFlags.Condition = `id eq ` + args[0]
		}
		client := initClient(cfg)
		target, err := cwctl.GetCommandsExecuted(client, paramsDefault.merge(&cliFlags), cliTargetID)
		if err != nil {
			Failf("error attempting GetComputerCommands: %v", err)
		}
		if len(target) < 1 {
			Infof("No Results Found.")
			return
		}
		switch {
		case cliFlags.Condition != "":
			switch {
			case cmd.Flags().Changed(`out`):
				handlePrint(target, outFormat)
			default:
				raw := []byte(target[0].Output)
				if len(raw) < 1 {
					handlePrint(target, `table`)
					return
				}
				handlePrint(raw, `raw`)
			}
		default:
			handlePrint(target, outFormat)
		}
	},
}

func init() {
	cmdComputerCommands.Flags().StringVarP(&cliTargetID, "computer-id", `C`, "", "ID of the Computer to target.")
	cmdComputerCommands.Flags().StringVarP(&cliFlags.Page, "page", `p`, cliFlags.Page, "Page number of results.")
	cmdComputerCommands.Flags().StringVarP(&cliFlags.PageSize, "page-size", `s`, cliFlags.PageSize, "Results per page.")
	cmdComputerCommands.MarkFlagRequired(`computer-id`)
}
