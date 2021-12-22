package cmd

import (
	"github.com/spf13/cobra"
)

var cmdRefreshToken = &cobra.Command{
	Use:     "refresh-token",
	Aliases: []string{`refresh`},
	Short:   "refresh existing access token",
	Run: func(cmd *cobra.Command, args []string) {
		client := initClient(cfg)
		handleRefresh(client)
	},
}
