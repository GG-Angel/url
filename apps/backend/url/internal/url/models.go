package url

import repo "github.com/GG-Angel/url/internal/adapters/postgresql/sqlc"

type URLWithTags struct {
	repo.Url
	Tags []repo.Tag `json:"tags"`
}

func NewURLWithTags(u repo.Url, t []repo.Tag) URLWithTags {
	return URLWithTags{Url: u, Tags: t}
}
