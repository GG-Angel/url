package url

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log/slog"
	"math/big"

	repo "github.com/GG-Angel/url/internal/adapters/postgresql/sqlc"
	"github.com/jackc/pgx/v5/pgconn"
)

type service interface {
	GetUrlByCode(ctx context.Context, code string) (repo.Url, error)
	ShortenUrl(ctx context.Context, url string) (repo.Url, error)
}

type svc struct {
	repo repo.Querier
}

// GetUrlByCode implements [service].
func (s *svc) GetUrlByCode(ctx context.Context, code string) (repo.Url, error) {
	return s.repo.FindUrlByCode(ctx, code)
}

// ShortenUrl implements [service].
func (s *svc) ShortenUrl(ctx context.Context, url string) (repo.Url, error) {
	const maxRetries = 5

	for attempt := range maxRetries {
		code := generateCode()

		url, err := s.repo.CreateUrl(ctx, repo.CreateUrlParams{Url: url, Code: code})
		if err == nil {
			slog.Info("created short url", "code", code)
			return url, nil
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			slog.Debug("code collision", "code", code, "attempt", attempt+1)
			continue // collision
		}

		slog.Error("failed to create short url", "error", err)
		return repo.Url{}, err
	}

	return repo.Url{}, fmt.Errorf("failed to generate code after %d attempts", maxRetries)
}

func generateCode() string {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const codeLength = 6

	buffer := make([]byte, codeLength)
	for i := range buffer {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		buffer[i] = alphabet[index.Int64()]
	}
	return string(buffer)
}

func NewService(repo repo.Querier) service {
	return &svc{repo}
}
