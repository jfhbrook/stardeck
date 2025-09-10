/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/godbus/dbus/v5"

	"github.com/jfhbrook/stardeck/client"
	"github.com/jfhbrook/stardeck/logger"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a value",
	Long: `Set an ephemeral value on the Stardeck service.`,
}

var setWindowCmd = &cobra.Command{
	Use:   "window [name]",
	Short: "Set the window title",
	Long: `Set the title of the currently active window. This will be displayed on
the LCD.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			logger.FlagrantError(errors.New("Window name is required"))
		}

		windowName := args[0]

		conn, err := dbus.ConnectSessionBus()

		if err != nil {
			logger.FlagrantError(err)
		}

		defer conn.Close()

		cl := client.NewStardeckClient(conn)

		cl.SetWindow(windowName)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.AddCommand(setWindowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
