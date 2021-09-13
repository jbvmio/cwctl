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
		var condition, ids string
		switch len(args) {
		case 0:
		case 1:
			ids = `id eq ` + args[0]
		default:
			ids = `id in(` + args[0]
			for _, id := range args[1:] {
				ids += `, ` + id
			}
			ids += `)`
		}
		switch {
		case !cmd.Flags().Changed(`client-id`):
			if len(args) < 1 {
				Failf("must use either --client-id and/or computer id(s) as arguments")
			}
			condition = ids
		default:
			switch len(args) {
			case 0:
				condition = `client.id eq ` + cliClientID
			default:
				condition = `(client.id eq ` + cliClientID + `) and ` + ids
			}
		}
		cliFlags.Condition = condition
		client := initClient(cfg)
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
}
