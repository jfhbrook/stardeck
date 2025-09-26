package lcd

import (
	"github.com/spf13/cobra"
)

var LcdCmd = &cobra.Command{
	Use:   "lcd",
	Short: "Set a property on the LCD",
	Long:  `Set a property on the LCD, such as the currently displayed screen.`,
}

func init() {
}
