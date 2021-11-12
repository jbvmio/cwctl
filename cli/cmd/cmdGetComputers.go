package cmd

import (
	"github.com/jbvmio/cwctl"
	"github.com/spf13/cobra"
)

var cmdGetComputers = &cobra.Command{
	Use:     "computer",
	Aliases: []string{"computers", "comps", "comp"},
	Short:   "get computer details",
	Run: func(cmd *cobra.Command, args []string) {
		var condition, ids string
		switch len(args) {
		case 0:
		case 1:
			client := initClient(cfg)
			target, err := cwctl.GetComputer(client, args[0])
			if err != nil {
				Failf("error attempting GetComputer: %v", err)
			}
			handlePrint([]cwctl.Computer{target}, outFormat)
			return
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
				condition = `client.id eq ` + cliTargetID
			default:
				condition = `(client.id eq ` + cliTargetID + `) and ` + ids
			}
		}
		cliFlags.Condition = condition
		client := initClient(cfg)
		target, err := cwctl.GetComputers(client, paramsComputer.merge(&cliFlags))
		if err != nil {
			Failf("error attempting GetComputers: %v", err)
		}
		handlePrint(target, outFormat)
	},
}

func init() {
	cmdGetComputers.Flags().StringVarP(&cliTargetID, "client-id", `C`, "", "ID of the Client to target.")
	cmdGetComputers.Flags().StringVarP(&cliFlags.Page, "page", `p`, cliFlags.Page, "Page number of results.")
	cmdGetComputers.Flags().StringVarP(&cliFlags.PageSize, "page-size", `s`, cliFlags.PageSize, "Results per page.")
	cmdGetComputers.AddCommand(cmdComputerCommands)
}
