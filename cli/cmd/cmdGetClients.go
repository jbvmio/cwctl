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
		target, err := cwctl.GetClients(client, paramsDefault.merge(&cliFlags))
		if err != nil {
			Failf("error attempting GetClients: %v", err)
		}
		handlePrint(target, outFormat)
	},
}

func init() {
	cmdGetClients.Flags().StringVarP(&cliFlags.Page, "page", `p`, cliFlags.Page, "Page number of results.")
	cmdGetClients.Flags().StringVarP(&cliFlags.PageSize, "page-size", `s`, cliFlags.PageSize, "Results per page.")
}
