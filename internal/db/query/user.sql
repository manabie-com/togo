-- name: CreateUser :one
INSERT INTO users (
    user_name,
    hashed_password,
    created_at,
    updated_at,
    maximum_task_in_day
) VALUES (
             $1, $2, $3, $4, $5
         ) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE user_name = $1 LIMIT 1;