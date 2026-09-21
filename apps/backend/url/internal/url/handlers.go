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
	url, err := h.service.GetUrlByCode(r.Context(), chi.URLParam(r, "code"))
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.Url, http.StatusPermanentRedirect)
}

func (h *handler) GetUrl(w http.ResponseWriter, r *http.Request) {
	id := middleware.GetID(r.Context())

	url, err := h.service.GetUrlByID(r.Context(), id)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	tags, err := h.service.GetTagsForUrl(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to get tags for URL", http.StatusInternalServerError)
		return
	}

	io.Write(w, http.StatusOK, UrlWithTagsFromRepo(url, tags))
}

func (h *handler) CreateUrl(w http.ResponseWriter, r *http.Request) {
	var req CreateUrlRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	url, err := h.service.ShortenUrl(r.Context(), req.Url, req.Tags, req.Slug)
	if err != nil {
		http.Error(w, "Failed to create URL", http.StatusInternalServerError)
		return
	}

	io.Write(w, http.StatusCreated, UrlFromRepo(url))
}

func (h *handler) ListUrls(w http.ResponseWriter, r *http.Request) {
	urls, err := h.service.ListUrls(r.Context())
	if err != nil {
		http.Error(w, "Failed to get URLs", http.StatusInternalServerError)
		return
	}

	response := make([]UrlWithTagsResponse, len(urls))
	for i, url := range urls {
		response[i] = UrlWithTagsFromRepo(url.Url, url.Tags)
	}

	io.Write(w, http.StatusOK, response)
}

func (h *handler) ListTags(w http.ResponseWriter, r *http.Request) {
	tags, err := h.service.ListTags(r.Context())
	if err != nil {
		http.Error(w, "Failed to get tags", http.StatusInternalServerError)
		return
	}

	io.Write(w, http.StatusOK, tags)
}

func (h *handler) DeleteUrl(w http.ResponseWriter, r *http.Request) {
	id := middleware.GetID(r.Context())

	if err := h.service.DeleteUrl(r.Context(), id); err != nil {
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
