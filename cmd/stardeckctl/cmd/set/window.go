package set

import (
	"errors"

	pkgerrors "github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/godbus/dbus/v5"

	"github.com/jfhbrook/stardeck/client"
	"github.com/jfhbrook/stardeck/logger"
)

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
		  logger.FlagrantError(pkgerrors.Wrap(err, "Failed to connect to session bus"))
		}

		defer conn.Close()

		cl := client.NewStardeckClient(conn)

		cl.SetWindow(windowName)
	},
}

func init() {
	SetCmd.AddCommand(setWindowCmd)
}
