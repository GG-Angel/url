-- name: GetURLByID :one
SELECT *
FROM urls
WHERE id = $1;

-- name: GetURLByCode :one
SELECT *
FROM urls
WHERE code = $1;

-- name: GetURLsWithTags :many
SELECT sqlc.embed(u), sqlc.embed(t)
FROM urls u
LEFT OUTER JOIN url_tags ut ON ut.url_id = u.id
LEFT OUTER JOIN tags t ON t.id = ut.tag_id
ORDER BY u.created_at DESC, t.id;

-- name: GetTags :many
SELECT *
FROM tags
ORDER BY name;

-- name: GetTagsForURL :many
SELECT t.*
FROM tags t
JOIN url_tags ut ON ut.tag_id = t.id
WHERE ut.url_id = $1
ORDER BY t.id;

-- name: GetVisits :many
SELECT *
FROM visits
ORDER BY visited_at DESC;

-- name: GetVisitsForURL :many
SELECT v.*
FROM visits v
WHERE v.url_id = $1
ORDER BY v.visited_at DESC;

-- name: CreateURL :one
INSERT INTO urls (url, code)
VALUES ($1, $2)
RETURNING *;

-- name: CreateTag :one
INSERT INTO tags (name)
VALUES ($1)
ON CONFLICT (name) DO UPDATE
    SET name = EXCLUDED.name
RETURNING *;

-- name: AddTagToURL :exec
INSERT INTO url_tags (url_id, tag_id)
VALUES ($1, $2);

-- name: RemoveTagFromURL :exec
DELETE FROM url_tags
WHERE url_id = $1 AND tag_id = $2;

-- name: DeleteURL :exec
DELETE FROM urls
WHERE id = $1;

-- name: DeleteTag :exec
DELETE FROM tags
WHERE id = $1;
