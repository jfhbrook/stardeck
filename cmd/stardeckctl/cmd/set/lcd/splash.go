package lcd

import (
	"github.com/godbus/dbus/v5"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/jfhbrook/stardeck/crystalfontz"
	"github.com/jfhbrook/stardeck/logger"
)

var splashCmd = &cobra.Command{
	Use:   "splash",
	Short: "Set the splash screen",
	Long: `Update the LCD to show the splash screen. Will get overwritten by the
Stardeck service, if it is running.`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := dbus.ConnectSystemBus()

		if err != nil {
			logger.FlagrantError(errors.Wrap(err, "Failed to connect to system bus"))
		}

		defer conn.Close()

		lcd := crystalfontz.NewClient(conn)

		err = lcd.Splash()

		if err != nil {
			logger.FlagrantError(err)
		}

		err = lcd.StoreBootState(crystalfontz.NilFloat, crystalfontz.NilInt)

		if err != nil {
			logger.FlagrantError(err)
		}

		log.Info().Msg("Set splash screen on LCD")
	},
}

func init() {
	LcdCmd.AddCommand(splashCmd)
}
