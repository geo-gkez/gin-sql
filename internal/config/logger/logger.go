package logger

import (
	"log/slog"
	"os"
	"strings"
)

var Logger *slog.Logger

// LoggerConfig holds the configuration for the logger
type LoggerConfig struct {
	Level string
}

func InitLogger(loggerConfig LoggerConfig) {
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: parseLogLevel(loggerConfig.Level),
	}))

	slog.SetDefault(Logger)
}

// parseLogLevel converts a string level to slog.Level
func parseLogLevel(level string) slog.Level {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN", "WARNING":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo // Default to Info if unrecognized
	}
}
