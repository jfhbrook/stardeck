package set

import (
	"github.com/spf13/cobra"

	"github.com/jfhbrook/stardeck/cmd/stardeckctl/cmd/set/lcd"
)

var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a value",
	Long:  `Set an ephemeral value on the Stardeck service.`,
}

func init() {
	SetCmd.AddCommand(lcd.LcdCmd)
}
