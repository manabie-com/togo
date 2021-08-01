-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = $1
LIMIT 1;

-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: InsertUser :one
INSERT INTO users (username, password)
VALUES (@username, @password)
RETURNING *;

