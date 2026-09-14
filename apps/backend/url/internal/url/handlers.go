package url

import (
	"encoding/json"
	"net/http"

	"github.com/GG-Angel/url/internal/io"
	"github.com/go-chi/chi/v5"
)

type handler struct {
	service service
}

func (h *handler) RedirectFromCode(w http.ResponseWriter, r *http.Request) {
	url, err := h.service.GetUrlByCode(r.Context(), chi.URLParam(r, "code"))
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.Url, http.StatusPermanentRedirect)
}

func (h *handler) GetUrlByCode(w http.ResponseWriter, r *http.Request) {
	url, err := h.service.GetUrlByCode(r.Context(), chi.URLParam(r, "code"))
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
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

	url, err := h.service.ShortenUrl(r.Context(), req.Url)
	if err != nil {
		http.Error(w, "Failed to create URL", http.StatusInternalServerError)
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
		http.Error(w, "Failed to list URLs", http.StatusInternalServerError)
		return
	}

	response := make([]UrlResponse, len(urls))
	for i, u := range urls {
		response[i] = UrlFromRepo(u)
	}
	io.Write(w, http.StatusOK, response)
}

func NewHandler(s service) *handler {
	return &handler{s}
}
