-- name: CreateUser :one
INSERT INTO users (
  username,
  password
) VALUES (
  $1, $2
)
RETURNING *;

-- name: RetrieveUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;