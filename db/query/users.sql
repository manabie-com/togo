-- name: CreateUser :one
INSERT INTO users (
  password
) VALUES (
  $1
)
RETURNING *;