package logs

import (
	"github.com/spf13/cobra"
)

var mopidyLogsCmd = &cobra.Command{
	Use:   "mopidy",
	Short: "Get mopidy logs",
	Long:  `Get the logs for the mopidy service, using journald.`,
	Run: func(cmd *cobra.Command, args []string) {
		journalctl("-u", "mopidy.service")
	},
}

func init() {
}
