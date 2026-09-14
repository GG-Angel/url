-- name: FindUrlByCode :one
SELECT * 
FROM urls 
WHERE code = $1;

-- name: FindUrlByID :one
SELECT * 
FROM urls 
WHERE id = $1;

-- name: ListUrls :many
SELECT *
FROM urls
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CreateUrl :one
INSERT INTO urls (url, code)
VALUES ($1, $2)
RETURNING *;
