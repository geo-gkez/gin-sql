package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

// LoggerConfig holds the configuration for the logger
type LoggerConfig struct {
	Level  slog.Level
	Format string
}

func InitLogger(loggerConfig LoggerConfig) {
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: loggerConfig.Level,
	}))

	slog.SetDefault(Logger)
}
