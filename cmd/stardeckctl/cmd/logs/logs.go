package logs

import (
	"github.com/spf13/cobra"
)

var LogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Fetch the logs for a component",
	Long: `Fetch the logs for a component. This command typically assumes the
logs are available through journalctl.`,
}

func init() {
	LogsCmd.AddCommand(crystalfontzLogsCmd)
	LogsCmd.AddCommand(kwinLogsCmd)
	LogsCmd.AddCommand(plusdeckLogsCmd)
	LogsCmd.AddCommand(stardeckLogsCmd)
}
