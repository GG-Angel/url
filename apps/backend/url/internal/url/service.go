package url

import (
	"context"
	"errors"

	repo "github.com/GG-Angel/url/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UrlWithTags struct {
	Url  repo.Url
	Tags []repo.Tag
}

type service interface {
	GetUrlByCode(ctx context.Context, code string) (repo.Url, error)
	GetUrlByID(ctx context.Context, id int) (repo.Url, error)
	GetTagsForUrl(ctx context.Context, urlID int) ([]repo.Tag, error)
	ListUrls(ctx context.Context, limit, offset int) ([]UrlWithTags, error)
	ShortenUrl(ctx context.Context, url string, tags []string) (repo.Url, error)
	DeleteUrl(ctx context.Context, id int) error
	DeleteTag(ctx context.Context, id int) error
}

var errInvalidLimitOrOffset = errors.New("invalid limit or offset")

type svc struct {
	repo *repo.Queries
	db   *pgxpool.Pool
}

// GetUrlByID implements [service].
func (s *svc) GetUrlByID(ctx context.Context, id int) (repo.Url, error) {
	return s.repo.GetUrlByID(ctx, int32(id))
}

// GetUrlByCode implements [service].
func (s *svc) GetUrlByCode(ctx context.Context, code string) (repo.Url, error) {
	return s.repo.GetUrlByCode(ctx, code)
}

// GetTagsForUrl implements [service].
func (s *svc) GetTagsForUrl(ctx context.Context, urlID int) ([]repo.Tag, error) {
	return s.repo.ListTagsForUrl(ctx, int32(urlID))
}

// ListUrls implements [service].
func (s *svc) ListUrls(ctx context.Context, limit int, offset int) ([]UrlWithTags, error) {
	if offset < 0 || limit <= 0 || limit > 100 {
		return nil, errInvalidLimitOrOffset
	}

	urls, err := s.repo.ListUrls(ctx, repo.ListUrlsParams{Limit: int32(limit), Offset: int32(offset)})
	if err != nil {
		return nil, err
	}

	result := make([]UrlWithTags, len(urls))
	for i, url := range urls {
		tags, err := s.repo.ListTagsForUrl(ctx, url.ID)
		if err != nil {
			return nil, err
		}
		result[i] = UrlWithTags{
			Url:  url,
			Tags: tags,
		}
	}

	return result, nil
}

// ShortenUrl implements [service].
func (s *svc) ShortenUrl(ctx context.Context, url string, tags []string) (repo.Url, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Url{}, err
	}
	defer tx.Rollback(ctx)
	qtx := s.repo.WithTx(tx)

	// create url with generated code
	urlRow, err := qtx.CreateUrl(ctx, repo.CreateUrlParams{Url: url, Code: generateCode()})
	if err != nil {
		return repo.Url{}, err
	}

	// create and attach tags to url
	for _, tag := range tags {
		tagRow, err := qtx.CreateTag(ctx, tag)
		if err != nil {
			return repo.Url{}, err
		}
		if err := qtx.AddTagToUrl(ctx, repo.AddTagToUrlParams{UrlID: urlRow.ID, TagID: tagRow.ID}); err != nil {
			return repo.Url{}, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return repo.Url{}, err
	}
	return urlRow, nil
}

// DeleteUrl implements [service].
func (s *svc) DeleteUrl(ctx context.Context, id int) error {
	return s.repo.DeleteUrl(ctx, int32(id))
}

// DeleteTag implements [service].
func (s *svc) DeleteTag(ctx context.Context, id int) error {
	return s.repo.DeleteTag(ctx, int32(id))
}

func NewService(repo *repo.Queries, db *pgxpool.Pool) service {
	return &svc{repo: repo, db: db}
}
