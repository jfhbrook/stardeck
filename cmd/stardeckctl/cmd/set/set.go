package set

import (
	"github.com/spf13/cobra"
)

var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a value",
	Long:  `Set an ephemeral value on the Stardeck service.`,
}

func init() {
}
