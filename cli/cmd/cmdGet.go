package cmd

import (
	"github.com/spf13/cobra"
)

var cmdGet = &cobra.Command{
	Use:   "get",
	Short: "get resource details",
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
	cmdGet.AddCommand(cmdGetClients)
}
