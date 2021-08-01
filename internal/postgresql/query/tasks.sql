-- name: ListTasks :many
SELECT *
FROM tasks
WHERE user_id = $1
ORDER BY id;

-- name: GetTask :one
SELECT *
FROM tasks
WHERE id = $1
  AND user_id = $2
LIMIT 1;

-- name: UpdateTask :exec
UPDATE tasks
SET is_done = $2
WHERE id = $1;

-- name: DeleteTask :exec
DELETE
FROM tasks
WHERE id = $1;
