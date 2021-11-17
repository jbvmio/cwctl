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
		target, err := cwctl.GetClients(client, paramsDefault.merge(&paramFlags))
		if err != nil {
			Failf("error attempting GetClients: %v", err)
		}
		handlePrint(target, outFormat)
	},
}

func init() {
	cmdGetClients.Flags().StringVarP(&paramFlags.Page, "page", `p`, paramFlags.Page, "Page number of results.")
	cmdGetClients.Flags().StringVarP(&paramFlags.PageSize, "page-size", `s`, paramFlags.PageSize, "Results per page.")
}
