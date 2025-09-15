package set

import (
	"github.com/godbus/dbus/v5"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/jfhbrook/stardeck/client"
	"github.com/jfhbrook/stardeck/logger"
	"github.com/jfhbrook/stardeck/loopback"
)

var enable bool
var disable bool
var manage bool
var noManage bool

var setLoopbackCmd = &cobra.Command{
	Use:   "loopback",
	Short: "Configure loopback settings",
	Long:  `Enable/disable or manage/unmanage audio loopback`,
	Run: func(cmd *cobra.Command, args []string) {
		if !(enable || disable || manage || noManage) {
			log.Warn().Msg("No actions taken")
			return
		}

		if enable || disable {
			manager := loopback.NewLoopbackManager("", -1, -1)

			if enable {
				manager.Enable()
				log.Info().Msg("Enabled loopback")
			} else {
				manager.Disable()
				log.Info().Msg("Disabled loopback")
			}
		}

		if manage || noManage {
			conn, err := dbus.ConnectSessionBus()

			if err != nil {
				logger.FlagrantError(errors.Wrap(err, "Failed to connect to Stardeck service"))
			}

			defer conn.Close()

			cl := client.NewClient(conn)

			if err := cl.SetLoopback(manage); err != nil {
				logger.FlagrantError(err)
			}

			if manage {
				log.Info().Msg("Loopback is now managed")
			} else {
				log.Info().Msg("Loopback is no longer managed")
			}
		}
	},
}

func init() {
	SetCmd.AddCommand(setLoopbackCmd)

	setLoopbackCmd.Flags().BoolVar(&enable, "enable", false, "Enable loopback")
	setLoopbackCmd.Flags().BoolVar(&disable, "disable", false, "Disable loopback")
	setLoopbackCmd.MarkFlagsMutuallyExclusive("enable", "disable")

	setLoopbackCmd.Flags().BoolVar(&manage, "manage", false, "Manage loopback")
	setLoopbackCmd.Flags().BoolVar(&noManage, "no-manage", false, "Do not manage loopback")
	setLoopbackCmd.MarkFlagsMutuallyExclusive("manage", "no-manage")
}
