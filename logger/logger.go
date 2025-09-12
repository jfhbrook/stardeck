package logger

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func getLogLevel(logLevel string) zerolog.Level {
	level, err := zerolog.ParseLevel(strings.ToLower(logLevel))

	if err != nil {
		log.Warn().Str("log-level", logLevel).Msg("Unknown log level")
		return zerolog.InfoLevel
	}

	return level
}

func colorize(formatted string, code int, disabled bool) string {
	if disabled {
		return formatted
	}

	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", code, formatted)
}

// TODO: --no-color flag
// TODO: Pull log level from config
func ConfigureLogger(logLevel string) {
	lvl := getLogLevel(logLevel)
	zerolog.SetGlobalLevel(lvl)

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	writer := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.DateTime,
	}

	writer.FormatLevel = func(i interface{}) string {
		level, _ := zerolog.ParseLevel(i.(string))
		formatted := fmt.Sprintf("%+5s:", i)
		return colorize(formatted, zerolog.LevelColors[level], false)
	}

	log.Logger = log.Output(writer)

}

func FlagrantError(err error) {
	log.Fatal().Stack().Err(err).Msg("Flagrant error!")
}
