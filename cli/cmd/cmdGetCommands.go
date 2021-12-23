package cmd

import (
	"github.com/jbvmio/cwctl"
	"github.com/spf13/cobra"
)

var cmdGetCommands = &cobra.Command{
	Use:     "command",
	Aliases: []string{"commands", "cmd", "cmds"},
	Short:   "list / get commands",
	Run: func(cmd *cobra.Command, args []string) {
		var condition string
		field := `name`
		op := `like`
		pct := `%`
		if cliExactMatch || len(cliTargets) > 0 {
			op = `eq`
			pct = ""
		}
		var terms []string
		switch len(cliTargets) {
		case 0:
			terms = args
		default:
			terms = cliTargets
			field = `id`
		}
		switch len(terms) {
		case 0:
		case 1:
			condition = field + ` ` + op + ` "` + pct + terms[0] + pct + `"`
		default:
			condition = field + ` ` + op + ` "` + pct + terms[0] + pct + `"`
			for _, a := range terms[1:] {
				condition += ` or ` + field + ` ` + op + ` "` + pct + a + pct + `"`
			}
		}
		paramFlags.Condition = condition
		client := initClient(cfg)
		target, err := cwctl.GetCommands(client, paramsDefault.merge(&paramFlags))
		if err != nil {
			Failf("error attempting GetClients: %v", err)
		}
		handlePrint(target, outFormat)
	},
}

func init() {
	cmdGetCommands.Flags().StringSliceVar(&cliTargets, "ids", cliTargets, "Target Command IDs.")
	cmdGetCommands.Flags().BoolVarP(&cliExactMatch, "exact", "x", false, "Exact Match.")
	cmdGetCommands.Flags().StringVarP(&paramFlags.Page, "page", `p`, paramFlags.Page, "Page number of results.")
	cmdGetCommands.Flags().StringVarP(&paramFlags.PageSize, "page-size", `s`, paramFlags.PageSize, "Results per page.")
}
