package cmd

import (
	"github.com/jbvmio/cwctl/connectwise"
	"github.com/spf13/cobra"
)

var (
	rawMethod string
	rawPath   string
	rawBody   string
	epID      int
	rawFlags  connectwise.Parameters
)

var cmdRaw = &cobra.Command{
	Use:   "raw",
	Short: "perform raw REST requests",
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case cmd.Flags().Changed(`id`):
			client := initClient(cfg)
			v := make([]interface{}, len(args))
			for i := 0; i < len(args); i++ {
				v[i] = args[i]
			}
			b, err := client.RawRestID(connectwise.EP(epID), rawBody, &rawFlags, v...)
			if err != nil {
				Failf("error making raw rest call with ID %d: %v", epID, err)
			}
			handlePrint(b, `raw`)
		case cmd.Flags().Changed(`path`):
			client := initClient(cfg)
			b, err := client.RawRestPath(args[0], rawMethod, rawBody, &rawFlags)
			if err != nil {
				Failf("error making raw rest call with ID %d: %v", epID, err)
			}
			handlePrint(b, `raw`)
		default:
			v := make([]interface{}, len(args))
			for i := 0; i < len(args); i++ {
				v[i] = args[i]
			}
			EPs := connectwise.GetEndPoints(v...)
			handlePrint(EPs, outFormat)
		}
	},
}

func init() {
	cmdRaw.Flags().IntVar(&epID, "id", -1, "ID of Desired Endpoint Path.")
	cmdRaw.Flags().StringVar(&rawPath, "path", "", "Endpoint Path.")
	cmdRaw.Flags().StringVarP(&rawMethod, "method", "m", "GET", "Desired Method.")
	cmdRaw.Flags().StringVarP(&rawBody, "body", "d", "", "Desired Body.")

	cmdRaw.Flags().StringVarP(&rawFlags.Page, `page`, `p`, "", "[Parameter] Use to page the results returned in the response.")
	cmdRaw.Flags().StringVarP(&rawFlags.PageSize, `pagesize`, `s`, "", "[Parameter] Controls the size of pages returned in the response.")
	cmdRaw.Flags().StringVar(&rawFlags.OrderBy, `orderby`, "", "[Parameter] Used to sort your results by a fieldname.")
	cmdRaw.Flags().StringVarP(&rawFlags.IncludeFields, `include-fields`, `I`, "", "[Parameter] Specifies a list of fields to be included in the response.")
	cmdRaw.Flags().StringVarP(&rawFlags.Ids, `ids`, `i`, "", "[Parameter] Used to specify a list of object ids.")
	cmdRaw.Flags().StringVarP(&rawFlags.Expand, `expand`, `e`, "", "[Parameter] Grabs entire child object instead of name and id.")
	cmdRaw.Flags().StringVarP(&rawFlags.ExcludeFields, `exclude-fields`, `E`, "", "[Parameter] Specifies a list of fields to be excluded in the response.")
	cmdRaw.Flags().StringVarP(&rawFlags.Condition, `condition`, `c`, "", "[Parameter] Used as search parameters to filter results.")
}
