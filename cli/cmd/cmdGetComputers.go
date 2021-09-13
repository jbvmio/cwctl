package cmd

import (
	"github.com/jbvmio/cwctl"
	"github.com/spf13/cobra"
)

var cmdGetComputers = &cobra.Command{
	Use:     "computers",
	Aliases: []string{"computer", "comps", "comp"},
	Short:   "get computer details",
	Run: func(cmd *cobra.Command, args []string) {
		client := initClient(cfg)
		cliFlags.Condition = `client.id eq ` + cliClientID
		computers, err := cwctl.GetComputers(client, paramsComputer.merge(&cliFlags))
		if err != nil {
			Failf("error attempting GetComputers: %v", err)
		}
		handlePrint(computers, outFormat)
	},
}

func init() {
	cmdGetComputers.Flags().StringVarP(&cliClientID, "client-id", `C`, "", "ID of the Client to target.")
	cmdGetComputers.Flags().StringVarP(&cliFlags.Page, "page", `p`, cliFlags.Page, "Page number of results.")
	cmdGetComputers.Flags().StringVarP(&cliFlags.PageSize, "page-size", `s`, cliFlags.PageSize, "Results per page.")
	cmdGetComputers.MarkFlagRequired(`client-id`)
}
