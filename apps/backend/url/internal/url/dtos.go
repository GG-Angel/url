package url

import (
	"time"

	repo "github.com/GG-Angel/url/internal/adapters/postgresql/sqlc"
)

type CreateUrlRequest struct {
	Url  string   `json:"url"`
	Tags []string `json:"tags"`
	Slug *string  `json:"slug,omitempty"`
}

type UrlResponse struct {
	ID        int        `json:"id"`
	Url       string     `json:"url"`
	Code      string     `json:"code"`
	CreatedAt time.Time  `json:"created_at"`
	Tags      []repo.Tag `json:"tags"`
}

func NewUrlResponse(url repo.Url, tags []repo.Tag) UrlResponse {
	return UrlResponse{
		ID:        int(url.ID),
		Url:       url.Url,
		Code:      url.Code,
		CreatedAt: url.CreatedAt,
		Tags:      tags,
	}
}
