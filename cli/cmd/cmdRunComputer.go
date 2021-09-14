package cmd

import (
	"github.com/spf13/cobra"
)

var cmdRunComputer = &cobra.Command{
	Use:     "computer",
	Aliases: []string{"computers", "comps", "comp"},
	Short:   "run / execute computer actions",
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			cmd.Help()
		default:
			Failf("no such resource: %q", args[0])
		}
	},
}

func init() {
	cmdRunComputer.AddCommand(cmdRunComputerCommands)
}
