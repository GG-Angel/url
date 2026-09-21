package url

import (
	"context"
	"log/slog"

	repo "github.com/GG-Angel/url/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type service interface {
	GetUrlByCode(ctx context.Context, code string) (repo.Url, error)
	GetUrlByID(ctx context.Context, id int) (UrlWithTags, error)
	ListUrls(ctx context.Context) ([]UrlWithTags, error)
	ListTags(ctx context.Context) ([]repo.ListTagsRow, error)
	ShortenUrl(ctx context.Context, url string, tags []string, slug *string) (repo.Url, error)
	DeleteUrl(ctx context.Context, id int) error
	DeleteTag(ctx context.Context, id int) error
}

type svc struct {
	repo *repo.Queries
	db   *pgxpool.Pool
}

// GetUrlByCode implements [service].
func (s *svc) GetUrlByCode(ctx context.Context, code string) (repo.Url, error) {
	return s.repo.GetUrlByCode(ctx, code)
}

// GetUrlByID implements [service].
func (s *svc) GetUrlByID(ctx context.Context, id int) (UrlWithTags, error) {
	url, err := s.repo.GetUrlByID(ctx, int32(id))
	if err != nil {
		return UrlWithTags{}, err
	}
	tags, err := s.repo.ListTagsForUrl(ctx, int32(id))
	if err != nil {
		return UrlWithTags{}, err
	}
	return UrlWithTags{
		ID:        int(url.ID),
		Url:       url.Url,
		Code:      url.Code,
		CreatedAt: url.CreatedAt,
		Tags:      tags,
	}, nil
}

// ListUrls implements [service].
func (s *svc) ListUrls(ctx context.Context) ([]UrlWithTags, error) {
	urls, err := s.repo.ListUrls(ctx)
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
			ID:        int(url.ID),
			Url:       url.Url,
			Code:      url.Code,
			CreatedAt: url.CreatedAt,
			Tags:      tags,
		}
	}

	return result, nil
}

// ListTags implements [service].
func (s *svc) ListTags(ctx context.Context) ([]repo.ListTagsRow, error) {
	return s.repo.ListTags(ctx)
}

// ShortenUrl implements [service].
func (s *svc) ShortenUrl(ctx context.Context, url string, tags []string, slug *string) (repo.Url, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return repo.Url{}, err
	}
	defer tx.Rollback(ctx)
	qtx := s.repo.WithTx(tx)

	// create url with generated code
	var code string
	if slug != nil {
		code = *slug
	} else {
		code = generateCode()
	}

	urlRow, err := qtx.CreateUrl(ctx, repo.CreateUrlParams{Url: url, Code: code})
	if err != nil {
		slog.Error("Failed to create URL", "error", err)
		return repo.Url{}, err
	}

	// create and attach tags to url
	for _, tag := range tags {
		tagRow, err := qtx.CreateTag(ctx, tag)
		if err != nil {
			slog.Error("Failed to create tag", "error", err)
			return repo.Url{}, err
		}
		if err := qtx.AddTagToUrl(ctx, repo.AddTagToUrlParams{UrlID: urlRow.ID, TagID: tagRow.ID}); err != nil {
			slog.Error("Failed to add tag to URL", "error", err)
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
