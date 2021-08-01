-- name: ListTasks :many
SELECT *
FROM tasks
ORDER BY id;

-- name: GetTask :one
SELECT *
FROM tasks
WHERE id = $1
LIMIT 1;

-- name: UpdateTask :exec
UPDATE tasks
SET is_done = $2
WHERE id = $1;
