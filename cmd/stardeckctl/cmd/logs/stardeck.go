package logs

import (
	"github.com/spf13/cobra"
)

var stardeckLogsCmd = &cobra.Command{
	Use:   "stardeck",
	Short: "Get stardeck logs",
	Long:  `Get the logs for the stardeck component, using journald.`,
	Run: func(cmd *cobra.Command, args []string) {
		journalctl("--user", "-u", "stardeck.service")
	},
}

func init() {
}
