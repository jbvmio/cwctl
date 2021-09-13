package cmd

import (
	"github.com/jbvmio/cwctl"
	"github.com/spf13/cobra"
)

var cmdGetClients = &cobra.Command{
	Use:     "clients",
	Aliases: []string{"client"},
	Short:   "get client details",
	Run: func(cmd *cobra.Command, args []string) {
		client := initClient(cfg)
		clients, err := cwctl.GetClients(client, &cwFlags)
		if err != nil {
			Failf("error attempting GetClients: %v", err)
		}
		handlePrint(clients, outFormat)
	},
}

func init() {
	//cmdGet.AddCommand(cmdCases)
}
