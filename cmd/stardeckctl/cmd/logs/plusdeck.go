package logs

import (
	"github.com/spf13/cobra"
)

var plusdeckLogsCmd = &cobra.Command{
	Use:   "plusdeck",
	Short: "Get plusdeck logs",
	Long:  `Get the logs for the plusdeck service, using journald.`,
	Run: func(cmd *cobra.Command, args []string) {
		journalctl("-u", "plusdeck.service")
	},
}

func init() {
}
