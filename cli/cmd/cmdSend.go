package cmd

import (
	"github.com/spf13/cobra"
)

var cmdSend = &cobra.Command{
	Use:   "send",
	Short: "send actions",
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
	cmdSend.AddCommand(cmdLocalFile)
}
