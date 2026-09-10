package internal

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateShortUrl(ctx context.Context, fullUrl string) (string, error) {
	const maxRetries = 5

	for range maxRetries {
		shortCode := generateCode()

		_, err := s.db.Exec(ctx, "INSERT INTO urls (short_code, full_url) VALUES ($1, $2)", shortCode, fullUrl)
		if err == nil {
			log.Println("successfully inserted short URL:", shortCode)
			return shortCode, nil
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			continue // collision, try again
		}

		log.Println("failed to insert short URL:", err)
		return "", err
	}

	log.Println("failed to generate short URL after maximum retries")
	return "", fmt.Errorf("failed to generate code after %d attempts", maxRetries)
}

func (s *Store) GetFullUrl(ctx context.Context, shortCode string) (string, error) {
	var fullUrl string
	err := s.db.QueryRow(ctx, "SELECT full_url FROM urls WHERE short_code = $1", shortCode).Scan(&fullUrl)
	return fullUrl, err
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
