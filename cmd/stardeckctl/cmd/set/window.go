package set

import (
	"github.com/godbus/dbus/v5"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
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
			logger.FlagrantError(errors.Wrap(err, "Failed to connect to Stardeck service"))
		}

		defer conn.Close()

		cl := client.NewClient(conn)

		if err := cl.SetWindow(windowName); err != nil {
			logger.FlagrantError(err)
		}

		log.Info().Str("name", windowName).Msg("Set window name")
	},
}

func init() {
	SetCmd.AddCommand(setWindowCmd)
}
