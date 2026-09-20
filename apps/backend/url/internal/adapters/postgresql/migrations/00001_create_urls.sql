-- +goose Up
CREATE TABLE IF NOT EXISTS urls (
    id SERIAL PRIMARY KEY,
    url TEXT NOT NULL,
    code TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX IF NOT EXISTS ix_urls_code 
ON urls (code);

-- +goose Down
DROP TABLE IF EXISTS urls;
