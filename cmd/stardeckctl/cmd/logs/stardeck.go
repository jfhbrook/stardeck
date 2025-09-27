package logs

import (
	"github.com/spf13/cobra"
)

var stardeckLogsCmd = &cobra.Command{
	Use:   "stardeck",
	Short: "Get stardeck logs",
	Long:  `Get the logs for the stardeck component, using journald.`,
	Run: func(cmd *cobra.Command, args []string) {
		argv := []string{"--user"}
		if boot {
			argv = append(argv, "-b")
		}
		if follow {
			argv = append(argv, "-f")
		}
		argv = append(argv, "-u", "stardeck.service")
		journalctl(argv...)
	},
}

func init() {
}
