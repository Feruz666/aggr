package logger

import (
	"os"

	"github.com/Feruz666/aggr/config"
	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

func NewLogger() *Logger {
	var logger Logger

	zeroLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	cfg, err := config.InitConfig()
	if err != nil {
		return nil
	}

	// Set proper loglevel based on config
	switch cfg.LogLevel {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn", "warning":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "err", "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel) // log info and above by default
	}

	logger = Logger{&zeroLogger}

	return &logger
}
