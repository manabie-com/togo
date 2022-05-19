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

-- name: SelectTaskByUserId :many
SELECT t.title, t.user_id, t.created_at, t.updated_at FROM tasks t
INNER JOIN  users u
ON t.user_id = u.id
WHERE u.id = $1 LIMIT 1;