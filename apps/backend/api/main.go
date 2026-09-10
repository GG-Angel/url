package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GG-Angel/url/api/internal"
	"github.com/jackc/pgx/v5/pgxpool"
)

func buildPostgresUrl() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
}

func main() {
	pool, err := pgxpool.New(context.Background(), buildPostgresUrl())
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	store := internal.NewStore(pool)
	app := internal.NewApp(store)

	http.HandleFunc("/health", app.HealthCheck)
	http.HandleFunc("/shorten", app.ShortenHandler)
	http.HandleFunc("/", app.RedirectHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
