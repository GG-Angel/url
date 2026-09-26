package url

import (
	"encoding/json"
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
	data, err := h.service.GetURLByCode(r.Context(), chi.URLParam(r, "code"))
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, data.Url.Url, http.StatusPermanentRedirect)
}

func (h *handler) GetURL(w http.ResponseWriter, r *http.Request) {
	id := middleware.GetID(r.Context())
	url, err := h.service.GetURLByID(r.Context(), id)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	io.Write(w, http.StatusOK, url)
}

func (h *handler) CreateURL(w http.ResponseWriter, r *http.Request) {
	var req CreateUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}
	url, err := h.service.ShortenURL(r.Context(), req.URL, req.Tags, req.Slug)
	if err != nil {
		http.Error(w, "Failed to create URL", http.StatusInternalServerError)
		return
	}
	io.Write(w, http.StatusCreated, url)
}

func (h *handler) GetURLs(w http.ResponseWriter, r *http.Request) {
	urls, err := h.service.GetURLs(r.Context())
	if err != nil {
		http.Error(w, "Failed to get URLs", http.StatusInternalServerError)
		return
	}
	io.Write(w, http.StatusOK, urls)
}

func (h *handler) GetTags(w http.ResponseWriter, r *http.Request) {
	tags, err := h.service.GetTags(r.Context())
	if err != nil {
		http.Error(w, "Failed to get tags", http.StatusInternalServerError)
		return
	}
	io.Write(w, http.StatusOK, tags)
}

func (h *handler) DeleteURL(w http.ResponseWriter, r *http.Request) {
	id := middleware.GetID(r.Context())
	if err := h.service.DeleteURL(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete URL", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *handler) DeleteTag(w http.ResponseWriter, r *http.Request) {
	id := middleware.GetID(r.Context())
	if err := h.service.DeleteTag(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete Tag", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func NewHandler(s service) *handler {
	return &handler{s}
}
