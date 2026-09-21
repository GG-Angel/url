package url

import (
	"time"

	repo "github.com/GG-Angel/url/internal/adapters/postgresql/sqlc"
)

type UrlWithTags struct {
	ID        int        `json:"id"`
	Url       string     `json:"url"`
	Code      string     `json:"code"`
	CreatedAt time.Time  `json:"created_at"`
	Tags      []repo.Tag `json:"tags"`
}
