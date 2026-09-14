-- name: FindUrlByCode :one
SELECT * 
FROM urls 
WHERE code = $1;

-- name: CreateUrl :one
INSERT INTO urls (url, code)
VALUES ($1, $2)
RETURNING *;
