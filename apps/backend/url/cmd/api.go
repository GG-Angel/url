package main

import (
	"log/slog"
	"net/http"
	"time"

	repo "github.com/GG-Angel/url/internal/adapters/postgresql/sqlc"
	appMiddleware "github.com/GG-Angel/url/internal/middleware"
	"github.com/GG-Angel/url/internal/url"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
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
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)

	urlService := url.NewService(repo.New(a.db), a.db)
	urlHandler := url.NewHandler(urlService)

	r.Get("/", urlHandler.HealthCheck)
	r.Get("/{code}", urlHandler.RedirectFromCode)

	// URL routes
	r.Route("/urls", func(r chi.Router) {
		r.Get("/", urlHandler.ListUrls)
		r.Post("/", urlHandler.CreateUrl)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(appMiddleware.ParseID)

			r.Get("/", urlHandler.GetUrl)
			r.Delete("/", urlHandler.DeleteUrl)
		})
	})

	// Tag routes
	r.Route("/tags", func(r chi.Router) {
		r.Get("/", urlHandler.ListTags)

		r.Route("/{id}", func(r chi.Router) {
			r.Use(appMiddleware.ParseID)

			r.Delete("/", urlHandler.DeleteTag)
		})
	})

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
