package main

import (
	"chess/internal/lib/logger"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	logger := logger.New()

	// storage, err := psql.New("")
	// if err != nil {
	// 	logger.Error("Error during database initialization", slog.String("err", err.Error()))
	// }

	r := chi.NewRouter()

	r.Route("/game", func(r chi.Router) {
		r.Get("/new", func(w http.ResponseWriter, r *http.Request) {

		})
	})

	logger.Info("Server started on port 8080")
	if err := http.ListenAndServe(":8000", r); err != nil {
		logger.Error(err.Error())
	}
}
