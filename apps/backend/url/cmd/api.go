package main

import (
	"log/slog"
	"net/http"
	"time"

	repo "github.com/GG-Angel/url/internal/adapters/postgresql/sqlc"
	"github.com/GG-Angel/url/internal/url"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

type application struct {
	config config
	db     *pgxpool.Pool
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}

func (a *application) mount() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	urlService := url.NewService(repo.New(a.db))
	urlHandler := url.NewHandler(urlService)
	r.Get("/{code}", urlHandler.RedirectFromCode)
	r.Get("/urls/{code}", urlHandler.GetUrlByCode)
	r.Post("/urls", urlHandler.CreateUrl)

	return r
}

func (a *application) run(h http.Handler) error {
	srv := http.Server{
		Addr:         a.config.addr,
		Handler:      h,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	slog.Info("server started", "addr", a.config.addr)
	return srv.ListenAndServe()
}
