-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT feeds.*, users.name AS user_name
FROM feeds
JOIN users
ON feeds.user_id = users.id;

-- name: GetFeedByURL :one
SELECT * FROM feeds
WHERE url = $1;