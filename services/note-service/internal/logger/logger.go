package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func New(isDevelopment bool) zerolog.Logger {
	var logger zerolog.Logger
	if isDevelopment {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	}
	return logger

}