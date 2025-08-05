package app

import (
	"chess/internal/config"
	"chess/internal/lib/logger"
	"chess/internal/storage"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
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
	mode, exists := os.LookupEnv("DEVELOPMENT")
	if !exists {
		return nil, errors.New("нужно указать переменную окружения DEVELOPMENT")
	}

	var isProd bool
	switch mode {
	case "dev":
		isProd = false
	case "prod":
		isProd = true
	default:
		return nil, fmt.Errorf("указана неправильная переменная окружения DEVELOPMENT, %s", mode)
	}

	config, err := config.New(isProd)
	if err != nil {
		return nil, err
	}

	log, err := logger.New(config.Logger, isProd)
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

	r.Route("/game", func(r chi.Router) {
		r.Post("/create", func(w http.ResponseWriter, r *http.Request) {

		})
	})

	srv := &http.Server{
		Addr:         app.config.GetAddr(),
		Handler:      r,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		app.log.Info("Сервер запущен", slog.Int("port", app.config.Port), slog.String("host", app.config.Host))
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
