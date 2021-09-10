package cmd

import (
	"github.com/spf13/cobra"
)

var cmdGetClients = &cobra.Command{
	Use:     "clients",
	Aliases: []string{"client"},
	Short:   "cwctl: get client details",
	Run: func(cmd *cobra.Command, args []string) {
		client := initClient(cfg)
		clients := getClients(client, &cwFlags)
		handlePrint(clients, outFormat)
	},
}

func init() {
	//cmdGet.AddCommand(cmdCases)
}
