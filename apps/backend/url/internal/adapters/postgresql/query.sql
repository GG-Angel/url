-- name: ListUrls :many
SELECT *
FROM urls
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListTagsForUrl :many
SELECT t.*
FROM tags t
JOIN url_tags ut ON ut.tag_id = t.id
WHERE ut.url_id = $1;

-- name: GetUrlByCode :one
SELECT * 
FROM urls 
WHERE code = $1;

-- name: GetUrlByID :one
SELECT * 
FROM urls 
WHERE id = $1;

-- name: CreateUrl :one
INSERT INTO urls (url, code)
VALUES ($1, $2)
RETURNING *;

-- name: CreateTag :one
INSERT INTO tags (name)
VALUES ($1)
ON CONFLICT (name) DO UPDATE
    SET name = EXCLUDED.name
RETURNING *;

-- name: AddTagToUrl :exec
INSERT INTO url_tags (url_id, tag_id)
VALUES ($1, $2);

-- name: RemoveTagFromUrl :exec
DELETE FROM url_tags
WHERE url_id = $1 AND tag_id = $2;

-- name: DeleteUrl :exec
DELETE FROM urls
WHERE id = $1;

-- name: DeleteTag :exec
DELETE FROM tags
WHERE id = $1;
