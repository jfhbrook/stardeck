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

const (
	JsonFormat   string = "json"
	PrettyFormat        = "pretty"
)

func getLogLevel(logLevel string) zerolog.Level {
	level, err := zerolog.ParseLevel(strings.ToLower(logLevel))

	if err != nil {
		log.Warn().Str("log-level", logLevel).Msg("Unknown log level")
		return zerolog.InfoLevel
	}

	return level
}

func colorize(formatted string, code int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", code, formatted)
}

func consoleWriter(color bool) zerolog.ConsoleWriter {
	writer := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.DateTime,
		NoColor:    !color,
	}

	writer.FormatLevel = func(i interface{}) string {
		level, _ := zerolog.ParseLevel(i.(string))
		formatted := fmt.Sprintf("%+5s:", i)

		if color {
			return colorize(formatted, zerolog.LevelColors[level])
		}
		return formatted
	}

	return writer
}

func ConfigureLogger(logLevel string, logFormat string, logColor bool) {
	lvl := getLogLevel(logLevel)
	zerolog.SetGlobalLevel(lvl)

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	if logFormat == PrettyFormat {
		log.Logger = log.Output(consoleWriter(logColor))
	}
}

func FlagrantError(err error) {
	log.Fatal().Stack().Err(err).Msg("Flagrant error!")
}
