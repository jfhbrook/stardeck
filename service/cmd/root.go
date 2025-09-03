package cmd

import (
	"os"

	"github.com/jfhbrook/stardeck/service/lib"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "service",
	Short: "Start the Stardeck service",
	Long: `stardeck-service is a service that manages capabilities specific to
the Stardeck 1A Media Appliance.`,
	Run: func(cmd *cobra.Command, args []string) {
		lib.Service()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
