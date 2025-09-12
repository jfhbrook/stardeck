package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jfhbrook/stardeck/config"
	"github.com/jfhbrook/stardeck/logger"
	"github.com/jfhbrook/stardeck/cmd/stardeckctl/cmd/get"
	"github.com/jfhbrook/stardeck/cmd/stardeckctl/cmd/set"
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

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(get.GetCmd)
	rootCmd.AddCommand(set.SetCmd)
}

func initConfig() {
	logger.ConfigureLogger()
	config.InitConfig()

	// TODO: https://github.com/spf13/cobra/blob/main/site/content/user_guide.md#create-rootcmd
	if err := viper.ReadInConfig(); err != nil {
		if !config.HandleConfigFileNotFoundError(err) {
			logger.FlagrantError(err)
		}
	}
}
