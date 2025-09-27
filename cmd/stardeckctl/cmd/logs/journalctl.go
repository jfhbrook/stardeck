package logs

import (
	"os"
	"os/exec"

	"github.com/jfhbrook/stardeck/logger"
)

func journalctl(args ...string) {
	cmd := exec.Command("journalctl", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		} else {
			logger.FlagrantError(err)
		}
	}
}
