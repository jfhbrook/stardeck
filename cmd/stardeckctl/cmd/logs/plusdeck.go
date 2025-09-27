package logs

import (
	"github.com/spf13/cobra"
)

var plusdeckLogsCmd = &cobra.Command{
	Use:   "plusdeck",
	Short: "Get plusdeck logs",
	Long:  `Get the logs for the plusdeck service, using journald.`,
	Run: func(cmd *cobra.Command, args []string) {
		argv := []string{}
		if boot {
			argv = append(argv, "-b")
		}
		if follow {
			argv = append(argv, "-f")
		}
		argv = append(argv, "-u", "plusdeck.service")
		journalctl(argv...)
	},
}

func init() {
}
