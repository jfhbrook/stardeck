package set

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

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
	SetCmd.AddCommand(setLoopbackCmd)

	setLoopbackCmd.Flags().BoolVar(&enable, "enable", false, "Enable loopback")
	setLoopbackCmd.Flags().BoolVar(&disable, "disable", false, "Disable loopback")
	setLoopbackCmd.MarkFlagsMutuallyExclusive("enable", "disable")

	setLoopbackCmd.Flags().BoolVar(&manage, "manage", false, "Manage loopback")
	setLoopbackCmd.Flags().BoolVar(&noManage, "no-manage", false, "Do not manage loopback")
	setLoopbackCmd.MarkFlagsMutuallyExclusive("manage", "no-manage")
}
