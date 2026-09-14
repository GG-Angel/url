package url

import (
	"time"

	repo "github.com/GG-Angel/url/internal/adapters/postgresql/sqlc"
)

type UrlResponse struct {
	Id        int       `json:"id"`
	Url       string    `json:"url"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateUrlRequest struct {
	Url string `json:"url"`
}

func UrlFromRepo(u repo.Url) UrlResponse {
	return UrlResponse{
		Id:        int(u.ID),
		Url:       u.Url,
		Code:      u.Code,
		CreatedAt: u.CreatedAt,
	}
}
