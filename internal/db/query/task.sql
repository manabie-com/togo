-- name: CreateTask :one
INSERT INTO tasks (
    title,
    user_id,
    created_at,
    updated_at
) VALUES (
             $1, $2, $3, $4
         ) RETURNING *;

-- name: GetTask :one
SELECT * FROM tasks
WHERE id = $1 LIMIT 1;

