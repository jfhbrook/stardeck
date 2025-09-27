package logs

import (
	"github.com/spf13/cobra"
)

var boot bool
var follow bool

var LogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Fetch the logs for a component",
	Long: `Fetch the logs for a component. This command typically assumes the
logs are available through journalctl.`,
}

func init() {
	LogsCmd.PersistentFlags().BoolVarP(&boot, "boot", "b", true, "Show logs since latest boot (default is 'true')")
	LogsCmd.PersistentFlags().BoolVarP(&follow, "follow", "f", false, "Tail the logs")

	LogsCmd.AddCommand(crystalfontzLogsCmd)
	LogsCmd.AddCommand(kwinLogsCmd)
	LogsCmd.AddCommand(mopidyLogsCmd)
	LogsCmd.AddCommand(plusdeckLogsCmd)
	LogsCmd.AddCommand(stardeckLogsCmd)
}
