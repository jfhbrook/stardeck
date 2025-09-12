package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func getLogLevel(logLevel string) zerolog.Level {
	lvl := strings.ToLower(logLevel)
	switch lvl {
	case "trace":
		return zerolog.TraceLevel
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	}
	log.Warn().Str("log-level", lvl).Msg("Unknown log level")
	return zerolog.InfoLevel
}

func ConfigureLogger(logLevel string) {
	lvl := getLogLevel(logLevel)
	zerolog.SetGlobalLevel(lvl)

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	// TODO: Custom writer
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func FlagrantError(err error) {
	log.Fatal().Stack().Err(err).Msg("Flagrant Error")
}
