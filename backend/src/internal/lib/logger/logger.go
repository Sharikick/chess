package logger

import (
	"chess/internal/config"
	"chess/internal/lib/buffer"
	"context"
	"encoding/json"
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

func New(config config.LoggerConfig, isProd bool) (*slog.Logger, error) {
	var handler slog.Handler
	if isProd {
		handler = &DevHandler{
			w: os.Stdout,
			opts: &DevHandlerOptions{
				level: config.Level,
			},
		}
	} else {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: config.Level,
		})
	}

	return slog.New(handler), nil
}
