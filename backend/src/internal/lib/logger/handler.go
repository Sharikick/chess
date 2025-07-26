package logger

import (
	"chess/internal/lib/buffer"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"sync"

	"github.com/fatih/color"
)

type DevHandler struct {
	w     io.Writer
	mutex sync.Mutex
}

func (h *DevHandler) Handle(_ context.Context, r slog.Record) error {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	buf := buffer.New()
	defer buf.Free()

	timeStr := r.Time.Format("[15:05:05.000]")
	buf.WriteString(timeStr)
	buf.WriteString(" ")

	level := r.Level.String() + ":"
	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	buf.WriteString(level)
	buf.WriteString(" ")

	msg := color.CyanString(r.Message)
	buf.WriteString(msg)
	buf.WriteString(" ")

	fields := make(map[string]any, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	buf.WriteString((string(b)))
	buf.WriteString("\n")

	_, err = h.w.Write(*buf)
	return err
}

// Проверяет будет ли обрабатываться лог с указанным уровнем логирования.
func (h *DevHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= slog.LevelDebug
}

func (h *DevHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *DevHandler) WithGroup(name string) slog.Handler {
	return h
}

func NewDevHandler(w io.Writer) slog.Handler {
	return &DevHandler{
		w: w,
	}
}
