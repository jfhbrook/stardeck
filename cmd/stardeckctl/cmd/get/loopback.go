package get

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/jfhbrook/stardeck/loopback"
)

var getLoopbackCmd = &cobra.Command{
	Use: "loopback",
	Short: "Get loopback settings",
	Long: `Get current settings for audio loopback`,
	Run: func(cmd *cobra.Command, args []string) {
		manager := loopback.NewLoopbackManager("", -1, -1)

		fmt.Println(manager.Status())
	},
}

func init() {
}
