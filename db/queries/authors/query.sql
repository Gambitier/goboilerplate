-- name: GetOne :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: List :many
SELECT * FROM authors
ORDER BY name;

-- name: Create :one
INSERT INTO authors (
  name, bio
) VALUES (
  $1, $2
)
RETURNING *;

-- name: Update :exec
UPDATE authors
  set name = $2,
  bio = $3
WHERE id = $1;

-- name: Delete :exec
DELETE FROM authors
WHERE id = $1;
