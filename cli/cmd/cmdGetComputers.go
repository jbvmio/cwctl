package cmd

import (
	"github.com/jbvmio/cwctl"
	"github.com/jbvmio/cwctl/connectwise"
	"github.com/spf13/cobra"
)

var cmdGetComputers = &cobra.Command{
	Use:     "computer",
	Aliases: []string{"computers", "comps", "comp"},
	Short:   "get computer details",
	Run: func(cmd *cobra.Command, args []string) {
		var condition string
		switch {
		case cliFlags.Query != "":
			condition = cliFlags.Query
		default:
			conditionals := &connectwise.Conditionals{}
			conditionals.AndIN("id", interfaceStrings(cliFlags.Targets)...).
				AndContains(connectwise.OR, "ComputerName", interfaceStrings(args)...)
			if cliTargetID != "" {
				conditionals.AndEquals(connectwise.OR, "client.id", cliTargetID)
			}
			condition = conditionals.String()
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
