package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jfhbrook/stardeck/cmd/stardeckctl/cmd/get"
	"github.com/jfhbrook/stardeck/cmd/stardeckctl/cmd/set"
	"github.com/jfhbrook/stardeck/config"
	"github.com/jfhbrook/stardeck/logger"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "stardeckctl",
	Short: "Control the Stardeck 1A Media Applicance",
	Long:  `Control the Stardeck 1A Media Appliance.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		config.InitConfig(cfgFile, config.Cli)
	})

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config-file", "", "Config file (default is /etc/stardeck/stardeck.yml)")
	rootCmd.PersistentFlags().String("log-level", "info", "Log level (default is 'info')")
	viper.BindPFlag("cli.log_level", rootCmd.PersistentFlags().Lookup("log-level"))

	rootCmd.PersistentFlags().String("log-format", logger.PrettyFormat, "Log format (default is 'pretty')")
	viper.BindPFlag("cli.log_format", rootCmd.PersistentFlags().Lookup("log-format"))

	rootCmd.PersistentFlags().Bool("log-color", true, "Show logs in color (default is 'true')")
	viper.BindPFlag("cli.log_color", rootCmd.PersistentFlags().Lookup("log-color"))

	rootCmd.AddCommand(get.GetCmd)
	rootCmd.AddCommand(set.SetCmd)
}
