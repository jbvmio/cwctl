package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	contactName string
	buildTime   string
	commitHash  string
)

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Print Version and Exit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("cwctl   : %s\n", contactName)
		fmt.Printf("version : %s\n", buildTime)
		fmt.Printf("commit  : %s\n", commitHash)
	},
}
