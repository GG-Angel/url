package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/GG-Angel/url/internal/env"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config{
		addr: ":8080",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "user=postgres password=postgres host=localhost dbname=url sslmode=disable"),
		},
	}

	// logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// database
	pool, err := pgxpool.New(context.Background(), cfg.db.dsn)
	if err != nil {
		slog.Error("failed to connect to database", "dsn", cfg.db.dsn, "error", err)
		os.Exit(1)
	}
	defer pool.Close()
	slog.Info("connected to database", "dsn", cfg.db.dsn)

	app := application{
		config: cfg,
		db:     pool,
	}

	if err := app.run(app.mount()); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
