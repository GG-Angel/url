-- +goose Up
CREATE TABLE IF NOT EXISTS urls (
    id SERIAL PRIMARY KEY,
    code TEXT NOT NULL,
    url TEXT NOT NULL,
    android_url TEXT,
    ios_url TEXT,
    desktop_url TEXT,
    enabled_since TIMESTAMPTZ,
    enabled_until TIMESTAMPTZ,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX IF NOT EXISTS ix_urls_code
ON urls (code);

-- +goose Down
DROP TABLE IF EXISTS urls;
