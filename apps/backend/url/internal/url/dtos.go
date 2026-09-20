package url

import (
	"time"

	repo "github.com/GG-Angel/url/internal/adapters/postgresql/sqlc"
)

type TagResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type UrlResponse struct {
	ID        int           `json:"id"`
	Url       string        `json:"url"`
	Code      string        `json:"code"`
	CreatedAt time.Time     `json:"created_at"`
	Tags      []TagResponse `json:"tags"`
}

type CreateUrlRequest struct {
	Url  string   `json:"url"`
	Tags []string `json:"tags"`
}

func UrlFromRepo(u repo.Url) UrlResponse {
	return UrlResponse{
		ID:        int(u.ID),
		Url:       u.Url,
		Code:      u.Code,
		CreatedAt: u.CreatedAt,
	}
}
