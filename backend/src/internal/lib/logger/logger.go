package logger

import (
	"chess/internal/config"
	"chess/internal/lib/buffer"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sync"

	"github.com/fatih/color"
)

type (
	DevHandlerOptions struct {
		level slog.Level
	}

	DevHandler struct {
		w     io.Writer
		mutex sync.Mutex

		opts *DevHandlerOptions
	}
)

func (handler *DevHandler) Handle(_ context.Context, record slog.Record) error {
	// Нужно для потокобезопасности
	handler.mutex.Lock()
	defer handler.mutex.Unlock()

	buf := buffer.New()
	defer buf.Free()

	buf.WriteString(record.Time.Format("[15:05:05.000]"))
	buf.WriteString(" ")

	var level string
	switch record.Level {
	case slog.LevelDebug:
		level = color.MagentaString("DBG:")
	case slog.LevelInfo:
		level = color.BlueString("INF:")
	case slog.LevelWarn:
		level = color.YellowString("WRN:")
	case slog.LevelError:
		level = color.RedString("ERR:")
	}

	buf.WriteString(level)
	buf.WriteString(" ")

	buf.WriteString(color.CyanString(record.Message))

	if record.NumAttrs() > 0 {
		buf.WriteString(" ")
		fields := make(map[string]any, record.NumAttrs())
		record.Attrs(func(a slog.Attr) bool {
			fields[a.Key] = a.Value.Any()
			return true
		})

		jsonStr, err := json.Marshal(fields)
		if err != nil {
			return err
		}

		buf.WriteString(string(jsonStr))
	}

	buf.WriteString("\n")

	_, err := handler.w.Write(*buf)
	return err
}

func (handler *DevHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= handler.opts.level
}

func (h *DevHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *DevHandler) WithGroup(name string) slog.Handler {
	return h
}

func parseLevel(level string) (slog.Level, error) {
	switch level {
	case "DEBUG":
		return slog.LevelDebug, nil
	case "INFO":
		return slog.LevelInfo, nil
	case "WARN":
		return slog.LevelWarn, nil
	case "ERROR":
		return slog.LevelError, nil
	default:
		return slog.LevelInfo, fmt.Errorf("Неправильно объявлен уровень лога: %s", level)
	}
}

func New(config config.LoggerConfig) (*slog.Logger, error) {
	level, err := parseLevel(config.Level)
	if err != nil {
		return nil, err
	}

	handler := &DevHandler{
		w: os.Stdout,
		opts: &DevHandlerOptions{
			level: level,
		},
	}

	return slog.New(handler), nil
}
