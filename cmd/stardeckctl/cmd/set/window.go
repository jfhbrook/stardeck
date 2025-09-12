package set

import (
	"github.com/godbus/dbus/v5"
	"github.com/spf13/cobra"

	"github.com/jfhbrook/stardeck/client"
)

var setWindowCmd = &cobra.Command{
	Use:   "window [name]",
	Short: "Set the window title",
	Long: `Set the title of the currently active window. This will be displayed on
the LCD.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		windowName := args[0]

		conn, err := dbus.ConnectSessionBus()

		if err != nil {
			return err
		}

		defer conn.Close()

		cl := client.NewStardeckClient(conn)

		cl.SetWindow(windowName)

		return nil
	},
}

func init() {
	SetCmd.AddCommand(setWindowCmd)
}
