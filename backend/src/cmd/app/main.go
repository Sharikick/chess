package main

import (
	"chess/internal/app"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app, err := app.New()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := app.Run(ctx); err != nil {
		log.Error("Ошибка при запуске приложения", slog.Any("error", err))
		os.Exit(1)
	}
}
