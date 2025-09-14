package set

import (
	"github.com/spf13/cobra"

	"github.com/jfhbrook/stardeck/client"
	"github.com/jfhbrook/stardeck/logger"
)

var setWindowCmd = &cobra.Command{
	Use:   "window [name]",
	Short: "Set the window title",
	Long: `Set the title of the currently active window. This will be displayed on
the LCD.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		windowName := args[0]

		cl, err := client.Connect()

		if err != nil {
			logger.FlagrantError(err)
		}

		cl.SetWindow(windowName)
	},
}

func init() {
	SetCmd.AddCommand(setWindowCmd)
}
