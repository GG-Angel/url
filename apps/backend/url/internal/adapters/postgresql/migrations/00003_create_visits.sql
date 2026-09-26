-- +goose Up
CREATE TABLE IF NOT EXISTS visits (
    id SERIAL PRIMARY KEY,
    url_id INT NOT NULL,
    ip_address INET NOT NULL,
    country TEXT NOT NULL,
    city TEXT NOT NULL,
    visited_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    FOREIGN KEY (url_id) REFERENCES urls(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS visits;
