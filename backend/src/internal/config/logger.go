package config

import (
	"fmt"
	"log/slog"
)

type LoggerConfig struct {
	Level slog.Level
}

func parseLevel(levelStr string) (slog.Level, error) {
	switch levelStr {
	case "DEBUG":
		return slog.LevelDebug, nil
	case "INFO":
		return slog.LevelInfo, nil
	case "WARN":
		return slog.LevelWarn, nil
	case "ERROR":
		return slog.LevelError, nil
	default:
		return slog.LevelInfo, fmt.Errorf("неправильно указан уровень для логов: %s (DEBUG, INFO, WARN, ERROR)", levelStr)
	}
}

func (l *LoggerConfig) fill(getEnv func(key, defaultValue string) (string, error)) error {
	levelStr, err := getEnv("LOG_LEVEL", "INFO")
	if err != nil {
		return err
	}
	level, err := parseLevel(levelStr)
	if err != nil {
		return err
	}
	l.Level = level

	return nil
}
