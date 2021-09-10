package cmd

import (
	"github.com/jbvmio/cwctl"
	"github.com/spf13/cobra"
)

var cwFlags cwctl.Parameters
var cmdGet = &cobra.Command{
	Use:   "get",
	Short: "cwctl: get resource details",
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			cmd.Help()
		default:
			Failf("no such resource: %q", args[0])
		}
	},
}

func init() {
	cmdGet.PersistentFlags().StringVarP(&cwFlags.Page, `page`, `p`, "", "Use to page the results returned in the response.")
	cmdGet.PersistentFlags().StringVarP(&cwFlags.PageSize, `pagesize`, `P`, "", "Controls the size of pages returned in the response.")
	cmdGet.PersistentFlags().StringVar(&cwFlags.OrderBy, `orderby`, "", "Used to sort your results by a fieldname.")
	cmdGet.PersistentFlags().StringVarP(&cwFlags.IncludeFields, `include-fields`, `I`, "", "Specifies a list of fields to be included in the response.")
	cmdGet.PersistentFlags().StringVarP(&cwFlags.Ids, `ids`, `i`, "", "Used to specify a list of object ids.")
	cmdGet.PersistentFlags().StringVarP(&cwFlags.Expand, `expand`, `e`, "", "Grabs entire child object instead of name and id.")
	cmdGet.PersistentFlags().StringVarP(&cwFlags.ExcludeFields, `exclude-fields`, `E`, "", "Specifies a list of fields to be excluded in the response.")
	cmdGet.PersistentFlags().StringVarP(&cwFlags.Condition, `condition`, `c`, "", "Used as search parameters to filter results.")
	cmdGet.AddCommand(cmdGetClients)
}
