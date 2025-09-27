package logs

import (
	"github.com/spf13/cobra"
)

var kwinLogsCmd = &cobra.Command{
	Use:   "kwin",
	Short: "Get KWin logs",
	Long:  `Get the logs for KWin, using journald.`,
	Run: func(cmd *cobra.Command, args []string) {
		journalctl("-b", "QT_CATEGORY=js", "QT_CATEGORY=kwin_scripting")
	},
}

func init() {
}
