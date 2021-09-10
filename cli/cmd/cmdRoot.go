package cmd

import (
	"fmt"
	"os"

	"github.com/jbvmio/cwctl"
	"github.com/spf13/cobra"
)

var (
	cfgFile   string
	outFormat string
	cfg       *Config
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cwctl",
	Short: "cwctl: ConnectWise CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute starts here.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", homeDir()+`/.cwctl.yaml`, "Path to config file")
	rootCmd.PersistentFlags().StringVarP(&outFormat, "out", "o", "", "Additional Output Formatting Options - raw|json|pretty|yaml.")
	rootCmd.AddCommand(cmdRaw)
}

func initConfig() {
	var err error
	cfg, err = GetConfig(cfgFile)
	if err != nil {
		Failf("error reading config: %v", err)
	}
}

func initClient(cfg *Config) *cwctl.Client {
	client, err := clientFromConfig(cfg)
	if err != nil {
		Failf("init err: %v", err)
	}
	return client
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
