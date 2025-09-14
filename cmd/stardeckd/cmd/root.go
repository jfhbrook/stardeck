package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/jfhbrook/stardeck/cmd/stardeckctl/cmd/get"
	"github.com/jfhbrook/stardeck/cmd/stardeckctl/cmd/set"
	"github.com/jfhbrook/stardeck/config"
	"github.com/jfhbrook/stardeck/logger"
	"github.com/jfhbrook/stardeck/service"
)

var (
	cfgFile string
)

var rootCmd = &cobra.Command{
	Use:   "stardeckd",
	Short: "Run the Stardeck 1A Media Applicance Service",
	Long:  `Run the Stardeck 1A Media Appliance Service.`,
	Run: func(cmd *cobra.Command, args []string) {
		service.Service()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		config.InitConfig(cfgFile, config.Service)
	})

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config-file", "", "Config file (default is /etc/stardeck/stardeck.yml)")
	rootCmd.PersistentFlags().String("log-level", "info", "Log level (default is 'info')")
	viper.BindPFlag("service.log_level", rootCmd.PersistentFlags().Lookup("log-level"))
	viper.SetDefault("service.log_level", "info")

	rootCmd.PersistentFlags().String("log-format", logger.JsonFormat, "Log format (default is 'json')")
	viper.BindPFlag("service.log_format", rootCmd.PersistentFlags().Lookup("log-format"))
	viper.SetDefault("service.log_format", logger.JsonFormat)

	rootCmd.PersistentFlags().Bool("log-color", false, "Show logs in color (default is 'false')")
	viper.BindPFlag("service.log_color", rootCmd.PersistentFlags().Lookup("log-color"))
	viper.SetDefault("service.log_color", false)

	rootCmd.AddCommand(get.GetCmd)
	rootCmd.AddCommand(set.SetCmd)
}
