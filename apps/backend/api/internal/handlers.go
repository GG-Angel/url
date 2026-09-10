package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type App struct {
	store *Store
}

func NewApp(store *Store) *App {
	return &App{
		store: store,
	}
}

func (a *App) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, ":3")
}

func (a *App) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	shortCode, err := a.store.CreateShortUrl(r.Context(), req.URL)
	if err != nil {
		http.Error(w, "Could not save url", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ShortenResponse{
		ShortCode: shortCode,
	})
}

func (a *App) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:]

	fullUrl, err := a.store.GetFullUrl(r.Context(), shortCode)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "URL not found", http.StatusNotFound)
		} else {
			http.Error(w, "Could not get url", http.StatusInternalServerError)
		}
		return
	}

	http.Redirect(w, r, fullUrl, http.StatusFound)
}
