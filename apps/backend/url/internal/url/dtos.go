package url

import (
	repo "github.com/GG-Angel/url/internal/adapters/postgresql/sqlc"
)

type CreateUrlRequest struct {
	URL  string   `json:"url"`
	Tags []string `json:"tags"`
	Slug *string  `json:"slug,omitempty"`
}

type UrlResponse struct {
	repo.Url
	Tags []repo.Tag `json:"tags"`
}

func NewUrlResponse(u repo.Url, t []repo.Tag) UrlResponse {
	return UrlResponse{u, t}
}
