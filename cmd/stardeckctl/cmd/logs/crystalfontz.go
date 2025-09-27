package logs

import (
	"github.com/spf13/cobra"
)

var crystalfontzLogsCmd = &cobra.Command{
	Use:   "crystalfontz",
	Short: "Get crystalfontz logs",
	Long:  `Get the logs for the crystalfontz service, using journald.`,
	Run: func(cmd *cobra.Command, args []string) {
		argv := []string{}
		if boot {
			argv = append(argv, "-b")
		}
		if follow {
			argv = append(argv, "-f")
		}
		argv = append(argv, "-u", "crystalfontz.service")
		journalctl(argv...)
	},
}

func init() {
}
