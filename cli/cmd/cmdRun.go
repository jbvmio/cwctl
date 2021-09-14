package cmd

import (
	"github.com/spf13/cobra"
)

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "run / execute actions",
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
	cmdRun.AddCommand(cmdRunComputer)
}
