package cmd

import (
	"github.com/jbvmio/cwctl"
	"github.com/spf13/cobra"
)

var cmdRelocate = &cobra.Command{
	Use:     "relocate",
	Aliases: []string{"send"},
	Short:   "relocate computers",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		body := cwctl.RelocatePayload{TargetId: cliFlags.Target, EntityIds: args}
		client := initClient(cfg)
		target, err := cwctl.RelocateComputers(client, paramsComputer.merge(&paramFlags), body)
		if err != nil {
			Failf("error attempting RelocateComputers: %v", err)
		}
		handlePrint(target, outFormat)
	},
}

func init() {
	cmdRelocate.Flags().StringVar(&cliFlags.Target, "location", cliFlags.Target, "Targeted Location ID.")
	cmdRelocate.Flags().StringSliceVarP(&cliFlags.Targets, "computer-id", `C`, cliFlags.Targets, "Targeted Computer IDs, comma delimited.")
	cmdRelocate.MarkFlagRequired("location")
}
