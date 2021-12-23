package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/jbvmio/cwctl/connectwise"
	"github.com/spf13/cobra"
)

var (
	cliFlags  CLIFlags
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
	rootCmd.PersistentFlags().StringVarP(&outFormat, "out", "o", "", "Additional Output Formatting Options - json|pretty|yaml.")
	rootCmd.AddCommand(cmdGet)
	rootCmd.AddCommand(cmdRun)
	rootCmd.AddCommand(cmdRaw)
	rootCmd.AddCommand(cmdSend)
	rootCmd.AddCommand(cmdRefreshToken)
	rootCmd.AddCommand(cmdVersion)
}

func initConfig() {
	var err error
	cfg, err = GetConfig(cfgFile)
	if err != nil {
		Failf("error reading config: %v", err)
	}
}

func initClient(cfg *Config) *connectwise.Client {
	client, err := clientFromConfig(cfg)
	if err != nil {
		Failf("init err: %v", err)
	}
	return client
}

func handleRefresh(client *connectwise.Client) {
	S, err := client.Token.SecondsLeft()
	if err != nil {
		Failf("error obtaining token seconds before expiry: %v", err)
	}
	switch {
	case S >= 3600:
		Infof("%v remaining on token...", time.Duration(S)*time.Second)
	case S >= 3300:
		Infof("%v remaining on prior token...", time.Duration(S)*time.Second)
	default:
		Infof("%v remaining on prior token... refreshing", time.Duration(S)*time.Second)
		err = client.RefreshToken()
		if err != nil {
			Failf("error refreshing token: %v", err)
		}
		err = client.SaveToken(cfg.TokenFile)
		if err != nil {
			Failf("error saving refreshed CW token: %v", err)
		}
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
