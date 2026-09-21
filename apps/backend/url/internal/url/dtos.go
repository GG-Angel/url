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
	ID        int       `json:"id"`
	Url       string    `json:"url"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}

type UrlWithTagsResponse struct {
	ID        int           `json:"id"`
	Url       string        `json:"url"`
	Code      string        `json:"code"`
	CreatedAt time.Time     `json:"created_at"`
	Tags      []TagResponse `json:"tags"`
}

type CreateUrlRequest struct {
	Url  string   `json:"url"`
	Tags []string `json:"tags"`
	Slug *string  `json:"slug,omitempty"`
}

func TagFromRepo(tag repo.Tag) TagResponse {
	return TagResponse{
		ID:        int(tag.ID),
		Name:      tag.Name,
		CreatedAt: tag.CreatedAt,
	}
}

func UrlFromRepo(url repo.Url) UrlResponse {
	return UrlResponse{
		ID:        int(url.ID),
		Url:       url.Url,
		Code:      url.Code,
		CreatedAt: url.CreatedAt,
	}
}

func UrlWithTagsFromRepo(url repo.Url, tags []repo.Tag) UrlWithTagsResponse {
	tagResponses := make([]TagResponse, len(tags))
	for i, tag := range tags {
		tagResponses[i] = TagFromRepo(tag)
	}

	return UrlWithTagsResponse{
		ID:        int(url.ID),
		Url:       url.Url,
		Code:      url.Code,
		CreatedAt: url.CreatedAt,
		Tags:      tagResponses,
	}
}
