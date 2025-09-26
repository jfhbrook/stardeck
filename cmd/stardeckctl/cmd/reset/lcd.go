package reset

import (
	"github.com/godbus/dbus/v5"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/jfhbrook/stardeck/crystalfontz"
	"github.com/jfhbrook/stardeck/logger"
)

var lcdCmd = &cobra.Command{
	Use:   "lcd",
	Short: "Reset the LCD",
	Long:  `Reset the LCD's brightness and contrast settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := dbus.ConnectSystemBus()

		if err != nil {
			logger.FlagrantError(errors.Wrap(err, "Failed to connect to system bus"))
		}

		defer conn.Close()

		lcd := crystalfontz.NewClient(conn)

		err = lcd.Reset()

		if err != nil {
			logger.FlagrantError(err)
		}

		log.Info().Msg("Reset settings on LCD")
	},
}

func init() {
	ResetCmd.AddCommand(lcdCmd)
}
