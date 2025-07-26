package logger

import (
	"log/slog"
	"os"
)

func New() *slog.Logger {
	return slog.New(NewDevHandler(os.Stdout))
}
