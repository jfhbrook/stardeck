package logs

import (
	"github.com/spf13/cobra"
)

var kwinLogsCmd = &cobra.Command{
	Use:   "kwin",
	Short: "Get KWin logs",
	Long:  `Get the logs for KWin, using journald.`,
	Run: func(cmd *cobra.Command, args []string) {
		argv := []string{}
		if boot {
			argv = append(argv, "-b")
		}
		if follow {
			argv = append(argv, "-f")
		}
		argv = append(argv, "QT_CATEGORY=js", "QT_CATEGORY=kwin_scripting")
		journalctl(argv...)
	},
}

func init() {
}
