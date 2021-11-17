package cmd

import (
	"strings"

	"github.com/jbvmio/cwctl"
	"github.com/spf13/cobra"
)

var cmdGetComputers = &cobra.Command{
	Use:     "computer",
	Aliases: []string{"computers", "comps", "comp"},
	Short:   "get computer details",
	Run: func(cmd *cobra.Command, args []string) {
		var condition, clientID, comps, ids string
		switch {
		case cliFlags.Query != "":
			condition = cliFlags.Query
		default:
			switch len(cliFlags.Targets) {
			case 0:
			case 1:
				client := initClient(cfg)
				target, err := cwctl.GetComputer(client, cliFlags.Targets[0])
				if err != nil {
					Failf("error attempting GetComputer: %v", err)
				}
				handlePrint([]cwctl.Computer{target}, outFormat)
				return
			default:
				ids = `id in(` + cliFlags.Targets[0]
				for _, id := range cliFlags.Targets[1:] {
					ids += `, ` + id
				}
				ids += `)`
			}
			switch len(args) {
			case 0:
			default:
				comps = `(ComputerName like '%` + args[0] + `%'`
				for _, cn := range args[1:] {
					comps += ` or ComputerName like '%` + cn + `%'`
				}
				comps += `)`
			}
			switch {
			case !cmd.Flags().Changed(`client-id`):
			default:
				clientID = `(client.id eq ` + cliTargetID + `)`
			}
			condition = clientID + ` and ` + comps + ` and ` + ids
			for strings.HasPrefix(condition, ` and `) {
				condition = strings.TrimPrefix(condition, ` and `)
			}
			for strings.HasSuffix(condition, ` and `) {
				condition = strings.TrimSuffix(condition, ` and `)
			}
		}
		paramFlags.Condition = condition
		client := initClient(cfg)
		target, err := cwctl.GetComputers(client, paramsComputer.merge(&paramFlags))
		if err != nil {
			Failf("error attempting GetComputers: %v", err)
		}
		handlePrint(target, outFormat)
	},
}

func init() {
	cmdGetComputers.Flags().StringVarP(&cliTargetID, "client-id", `C`, "", "ID of the Client to target.")
	cmdGetComputers.Flags().StringSliceVar(&cliFlags.Targets, "ids", cliFlags.Targets, "Desired Computer IDs.")
	cmdGetComputers.Flags().StringVarP(&cliFlags.Query, "query", `q`, "", "Use a Query.")
	cmdGetComputers.Flags().StringVarP(&paramFlags.Page, "page", `p`, paramFlags.Page, "Page number of results.")
	cmdGetComputers.Flags().StringVarP(&paramFlags.PageSize, "page-size", `s`, paramFlags.PageSize, "Results per page.")
	cmdGetComputers.AddCommand(cmdComputerCommands)
}
