-- +goose Up
CREATE TABLE IF NOT EXISTS tags (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS url_tags (
    id SERIAL PRIMARY KEY,
    url_id INT NOT NULL REFERENCES urls(id) ON DELETE CASCADE,
    tag_id INT NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (url_id, tag_id)
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_tags_name
ON tags (name);

-- +goose Down
DROP TABLE IF EXISTS url_tags;
DROP TABLE IF EXISTS tags;
