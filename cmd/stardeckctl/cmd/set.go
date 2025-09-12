/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"errors"

	"github.com/spf13/cobra"
	"github.com/godbus/dbus/v5"

	"github.com/jfhbrook/stardeck/client"
	"github.com/jfhbrook/stardeck/logger"
	"github.com/jfhbrook/stardeck/loopback"
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

var enable bool
var disable bool
var manage bool
var noManage bool

var setLoopbackCmd = &cobra.Command{
	Use: "loopback",
	Short: "Configure loopback settings",
	Long: `Enable/disable or manage/unmanage audio loopback`,
	Run: func(cmd *cobra.Command, args []string) {
		if enable && disable {
			logger.FlagrantError(errors.New("Can not both enable and disable loopback"))
		}

		if manage && noManage {
			logger.FlagrantError(errors.New("Can not both manage and not manage loopback"))
		}

		if enable || disable {
			manager := loopback.NewLoopbackManager("", "", -1, -1)

			if enable {
				manager.Enable()
			} else {
				manager.Disable()
			}
		}

		if manage || noManage {
			fmt.Println(manage)
			fmt.Println(noManage)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.AddCommand(setWindowCmd)
	setCmd.AddCommand(setLoopbackCmd)

	setLoopbackCmd.Flags().BoolVar(&enable, "enable", false, "Enable loopback")
	setLoopbackCmd.Flags().BoolVar(&disable, "disable", false, "Disable loopback")
	setLoopbackCmd.Flags().BoolVar(&manage, "manage", false, "Manage loopback")
	setLoopbackCmd.Flags().BoolVar(&noManage, "no-manage", false, "Do not manage loopback")
}
