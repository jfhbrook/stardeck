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
	RunE: func(cmd *cobra.Command, args []string) error {
		manager := loopback.NewLoopbackManager("", -1, -1)

		status, err := manager.Status()

		if err != nil {
			return err
		}

		fmt.Println(status)

		return nil
	},
}

func init() {
}
