package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pauljomy/social/internal/store"
)

type application struct {
	logger *slog.Logger
	config config
	posts  store.PostStorer
}

type config struct {
	addr string
	env  string
	db   dbConfig
}

type dbConfig struct {
	dsn          string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() http.Handler {
	mux := chi.NewRouter()

	mux.Use(app.logRequest)
	mux.Use(middleware.Recoverer)
	mux.Use(commonHeaders)
	mux.Use(enableCORS)

	mux.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthHandler)
	})

	return mux
}

func (app *application) run(mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	return srv.ListenAndServe()

}
