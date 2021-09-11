package cmd

import (
	"github.com/jbvmio/cwctl"
	"github.com/spf13/cobra"
)

var (
	rawMethod string
	rawBody   string
	epID      int
	cwFlags   cwctl.Parameters
)

var cmdRaw = &cobra.Command{
	Use:   "raw",
	Short: "perform raw REST requests",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case cmd.Flags().Changed(`id`):
			client := initClient(cfg)
			b, err := client.RawRestID(cwctl.EP(epID), rawBody, &cwFlags)
			if err != nil {
				Failf("error making raw rest call with ID %d: %v", epID, err)
			}
			handlePrint(b, `raw`)
		case len(args) > 0:
			client := initClient(cfg)
			b, err := client.RawRestPath(args[0], rawMethod, rawBody, &cwFlags)
			if err != nil {
				Failf("error making raw rest call with ID %d: %v", epID, err)
			}
			handlePrint(b, `raw`)
		default:
			EPs := cwctl.GetEndPoints()
			handlePrint(EPs, outFormat)

		}
	},
}

func init() {
	cmdRaw.Flags().IntVar(&epID, "id", -1, "ID of Desired Endpoint Path.")
	cmdRaw.Flags().StringVarP(&rawMethod, "method", "m", "GET", "Desired Method.")
	cmdRaw.Flags().StringVarP(&rawBody, "body", "d", "", "Desired Body.")

	cmdRaw.PersistentFlags().StringVarP(&cwFlags.Page, `page`, `p`, "", "[Parameter] Use to page the results returned in the response.")
	cmdRaw.PersistentFlags().StringVarP(&cwFlags.PageSize, `pagesize`, `P`, "", "[Parameter] Controls the size of pages returned in the response.")
	cmdRaw.PersistentFlags().StringVar(&cwFlags.OrderBy, `orderby`, "", "[Parameter] Used to sort your results by a fieldname.")
	cmdRaw.PersistentFlags().StringVarP(&cwFlags.IncludeFields, `include-fields`, `I`, "", "[Parameter] Specifies a list of fields to be included in the response.")
	cmdRaw.PersistentFlags().StringVarP(&cwFlags.Ids, `ids`, `i`, "", "[Parameter] Used to specify a list of object ids.")
	cmdRaw.PersistentFlags().StringVarP(&cwFlags.Expand, `expand`, `e`, "", "[Parameter] Grabs entire child object instead of name and id.")
	cmdRaw.PersistentFlags().StringVarP(&cwFlags.ExcludeFields, `exclude-fields`, `E`, "", "[Parameter] Specifies a list of fields to be excluded in the response.")
	cmdRaw.PersistentFlags().StringVarP(&cwFlags.Condition, `condition`, `c`, "", "[Parameter] Used as search parameters to filter results.")
}
