package logs

import (
	"github.com/spf13/cobra"
)

var crystalfontzLogsCmd = &cobra.Command{
	Use:   "crystalfontz",
	Short: "Get crystalfontz logs",
	Long:  `Get the logs for the crystalfontz service, using journald.`,
	Run: func(cmd *cobra.Command, args []string) {
		journalctl("-u", "crystalfontz.service")
	},
}

func init() {
}
