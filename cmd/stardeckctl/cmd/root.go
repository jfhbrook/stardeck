package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/jfhbrook/stardeck/cmd/stardeckctl/cmd/get"
	"github.com/jfhbrook/stardeck/cmd/stardeckctl/cmd/set"
	"github.com/jfhbrook/stardeck/config"
	"github.com/jfhbrook/stardeck/logger"
)

var rootCmd = &cobra.Command{
	Use:   "stardeckctl",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	cfgFile  string
	logLevel string
)

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config-file", "", "Config file (default is /etc/stardeck/stardeck.yml)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "Log level (default is 'info')")

	rootCmd.AddCommand(get.GetCmd)
	rootCmd.AddCommand(set.SetCmd)
}

func initConfig() {
	config.InitConfig(cfgFile)
	logger.ConfigureLogger(logLevel)
}
