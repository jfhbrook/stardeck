package reset

import (
	"github.com/spf13/cobra"
)

var ResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset a component",
	Long:  `Reset a component, such as the LCD`,
}

func init() {
}
