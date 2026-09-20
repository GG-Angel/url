package url

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/GG-Angel/url/internal/io"
	"github.com/GG-Angel/url/internal/middleware"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	service service
}

func (h *handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(":3"))
}

func (h *handler) RedirectFromCode(w http.ResponseWriter, r *http.Request) {
	url, err := h.service.GetUrlByCode(r.Context(), chi.URLParam(r, "code"))
	if err != nil {
		message := "URL not found"
		slog.Error(message, "error", err)
		http.Error(w, message, http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.Url, http.StatusPermanentRedirect)
}

func (h *handler) GetUrl(w http.ResponseWriter, r *http.Request) {
	id := middleware.GetID(r.Context())

	url, err := h.service.GetUrlByID(r.Context(), id)
	if err != nil {
		message := "URL not found"
		slog.Error(message, "error", err)
		http.Error(w, message, http.StatusNotFound)
		return
	}

	io.Write(w, http.StatusOK, UrlFromRepo(url))
}

func (h *handler) CreateUrl(w http.ResponseWriter, r *http.Request) {
	var req CreateUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	url, err := h.service.ShortenUrl(r.Context(), req.Url, req.Tags)
	if err != nil {
		message := "Failed to create URL"
		slog.Error(message, "error", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	io.Write(w, http.StatusCreated, UrlFromRepo(url))
}

func (h *handler) ListUrls(w http.ResponseWriter, r *http.Request) {
	limit := io.QueryParamInt(r, "limit", 10)
	offset := io.QueryParamInt(r, "offset", 0)
	if offset < 0 || limit <= 0 || limit > 100 {
		http.Error(w, "Invalid pagination parameters", http.StatusBadRequest)
		return
	}

	urls, err := h.service.ListUrls(r.Context(), limit, offset)
	if err != nil {
		message := "Failed to list URLs"
		slog.Error(message, "error", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	response := make([]UrlResponse, len(urls))
	for i, u := range urls {
		response[i] = UrlFromRepo(u)
	}
	io.Write(w, http.StatusOK, response)
}

func (h *handler) DeleteUrl(w http.ResponseWriter, r *http.Request) {
	id := middleware.GetID(r.Context())

	if err := h.service.DeleteUrl(r.Context(), id); err != nil {
		message := "Failed to delete URL"
		slog.Error(message, "error", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) DeleteTag(w http.ResponseWriter, r *http.Request) {
	id := middleware.GetID(r.Context())

	if err := h.service.DeleteTag(r.Context(), id); err != nil {
		message := "Failed to delete Tag"
		slog.Error(message, "error", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func NewHandler(s service) *handler {
	return &handler{s}
}
