package app

import (
	"chess/internal/config"
	"chess/internal/lib/logger"
	"chess/internal/storage"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	Application struct {
		config config.HTTPConfig
		log    *slog.Logger
		poolDB *pgxpool.Pool
	}
)

func New() (*Application, error) {
	config, err := config.New()
	if err != nil {
		return nil, err
	}

	log, err := logger.New(config.Logger)
	if err != nil {
		return nil, err
	}

	poolDB, err := storage.New(config.Database)
	if err != nil {
		return nil, err
	}

	return &Application{
		log:    log,
		config: config.HTTP,
		poolDB: poolDB,
	}, nil
}

func (app *Application) Run(ctx context.Context) error {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})

	srv := &http.Server{
		Addr:         app.config.Host + ":" + app.config.Port,
		Handler:      r,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		app.log.Info("Сервер запущен", slog.String("port", app.config.Port), slog.String("host", app.config.Host))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- fmt.Errorf("ошибка при запуске сервера: %w", err)
		}
		close(errCh)
	}()

	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		app.log.Info("Запуск изящной остановки приложения")
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			return err
		}
	}
	return nil
}
