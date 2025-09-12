package set

import (
	"github.com/godbus/dbus/v5"
	"github.com/spf13/cobra"

	"github.com/jfhbrook/stardeck/client"
	"github.com/jfhbrook/stardeck/logger"
)

var setWindowCmd = &cobra.Command{
	Use:   "window [name]",
	Short: "Set the window title",
	Long: `Set the title of the currently active window. This will be displayed on
the LCD.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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
	SetCmd.AddCommand(setWindowCmd)
}
