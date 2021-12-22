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
				AndContains(connectwise.OR, "ComputerName", interfaceStrings(args)...).
				AndContains(connectwise.OR, "OperatingSystemName", interfaceStrings(cliFlags.OSTargets)...)
			if cliFlags.Target != "" {
				conditionals.AndEquals(connectwise.OR, "client.id", cliFlags.Target)
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
	cmdGetComputers.Flags().StringVar(&cliFlags.Target, "client", cliFlags.Target, "Targeted Client ID.")
	cmdGetComputers.Flags().StringSliceVarP(&cliFlags.Targets, "computer-id", `C`, cliFlags.Targets, "Targeted Computer IDs, comma delimited.")
	cmdGetComputers.Flags().StringSliceVar(&cliFlags.OSTargets, "os", cliFlags.OSTargets, "Filter by OS, comma delimited.")
	cmdGetComputers.Flags().StringVarP(&cliFlags.Query, "query", `q`, cliFlags.Query, "Use a Query, Takes Precedent.")
	cmdGetComputers.Flags().StringVarP(&paramFlags.Page, "page", `p`, paramFlags.Page, "Page number of results.")
	cmdGetComputers.Flags().StringVarP(&paramFlags.PageSize, "page-size", `s`, paramFlags.PageSize, "Results per page.")
}
