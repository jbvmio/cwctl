package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rawMethod string
	rawBody   string
)

var cmdRaw = &cobra.Command{
	Use:   "raw",
	Short: "cwctl: Raw REST",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client := initClient(cfg)
		client.RawRest(args[0], rawMethod, rawBody)
	},
}

func init() {
	cmdRaw.Flags().StringVarP(&rawMethod, "method", "m", "GET", "Desired Method")
	cmdRaw.Flags().StringVarP(&rawBody, "body", "d", "", "Desired Body")
}
