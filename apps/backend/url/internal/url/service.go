package url

import (
	"context"
	"crypto/rand"
	"log/slog"
	"math/big"

	repo "github.com/GG-Angel/url/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type service interface {
	GetURLByCode(ctx context.Context, code string) (URLWithTags, error)
	GetURLByID(ctx context.Context, id int) (URLWithTags, error)

	GetURLs(ctx context.Context) ([]URLWithTags, error)
	GetTags(ctx context.Context) ([]repo.Tag, error)

	ShortenURL(ctx context.Context, url string, tags []string, slug *string) (URLWithTags, error)
	DeleteURL(ctx context.Context, id int) error
	DeleteTag(ctx context.Context, id int) error
}

type serviceImpl struct {
	repo *repo.Queries
	db   *pgxpool.Pool
}

func NewService(repo *repo.Queries, db *pgxpool.Pool) service {
	return &serviceImpl{repo: repo, db: db}
}

// GetTags implements [service].
func (s *serviceImpl) GetTags(ctx context.Context) ([]repo.Tag, error) {
	return s.repo.GetTags(ctx)
}

// GetURLByCode implements [service].
func (s *serviceImpl) GetURLByCode(ctx context.Context, code string) (URLWithTags, error) {
	url, err := s.repo.GetURLByCode(ctx, code)
	if err != nil {
		return URLWithTags{}, err
	}
	return s.attachTagsToUrl(ctx, url)
}

// GetURLByID implements [service].
func (s *serviceImpl) GetURLByID(ctx context.Context, id int) (URLWithTags, error) {
	url, err := s.repo.GetURLByID(ctx, int32(id))
	if err != nil {
		return URLWithTags{}, err
	}
	return s.attachTagsToUrl(ctx, url)

// GetURLs implements [service].
func (s *serviceImpl) GetURLs(ctx context.Context) ([]URLWithTags, error) {
	rows, err := s.repo.GetURLsWithTags(ctx)
	if err != nil {
		return []URLWithTags{}, err
	}

	urls := make([]URLWithTags, 0, len(rows))
	urlIndexByID := make(map[int32]int, len(rows))
	for _, row := range rows {
		idx, exists := urlIndexByID[row.Url.ID]
		if !exists {
			idx = len(urls)
			urls = append(urls, NewURLWithTags(row.Url, []repo.Tag{}))
			urlIndexByID[row.Url.ID] = idx
		}
		if row.Tag.Name != "" {
			urls[idx].Tags = append(urls[idx].Tags, row.Tag)
		}
	}
	return urls, nil
}

// ShortenURL implements [service].
func (s *serviceImpl) ShortenURL(ctx context.Context, url string, tags []string, slug *string) (URLWithTags, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return URLWithTags{}, err
	}
	defer tx.Rollback(ctx)
	qtx := s.repo.WithTx(tx)

	// create url with generated code
	var code string
	if slug != nil {
		code = *slug
	} else {
		code = s.generateCode()
	}

	urlRow, err := qtx.CreateURL(ctx, repo.CreateURLParams{Url: url, Code: code})
	if err != nil {
		slog.Error("Failed to create URL", "error", err)
		return URLWithTags{}, err
	}

	// create and attach tags to url
	for _, tag := range tags {
		tagRow, err := qtx.CreateTag(ctx, tag)
		if err != nil {
			slog.Error("Failed to create tag", "error", err)
			return URLWithTags{}, err
		}
		if err := qtx.AddTagToURL(ctx, repo.AddTagToURLParams{UrlID: urlRow.ID, TagID: tagRow.ID}); err != nil {
			slog.Error("Failed to add tag to URL", "error", err)
			return URLWithTags{}, err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return URLWithTags{}, err
	}
	return s.attachTagsToUrl(ctx, urlRow)
}

// DeleteURL implements [service].
func (s *serviceImpl) DeleteURL(ctx context.Context, id int) error {
	return s.repo.DeleteURL(ctx, int32(id))
}

// DeleteTag implements [service].
func (s *serviceImpl) DeleteTag(ctx context.Context, id int) error {
	return s.repo.DeleteTag(ctx, int32(id))
}

func (s *serviceImpl) attachTagsToUrl(ctx context.Context, url repo.Url) (URLWithTags, error) {
	tags, err := s.repo.GetTagsForURL(ctx, url.ID)
	if err != nil {
		return URLWithTags{}, err
	}
	return NewURLWithTags(url, tags), nil
}

func (s *serviceImpl) generateCode() string {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const codeLength = 6

	buffer := make([]byte, codeLength)
	for i := range buffer {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		buffer[i] = alphabet[index.Int64()]
	}
	return string(buffer)
}
