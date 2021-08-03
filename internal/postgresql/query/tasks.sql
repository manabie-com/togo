-- name: ListTasks :many
SELECT *
FROM tasks
WHERE user_id = @user_id
  AND (@created_date::date = '0001-01-01' OR created_date = @created_date)
  AND (NOT @is_done::boolean OR is_done = @is_done)
ORDER BY id;

-- name: GetTask :one
SELECT *
FROM tasks
WHERE id = $1
  AND user_id = $2
LIMIT 1;

-- name: UpdateTask :one
UPDATE tasks
SET is_done    = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteTask :exec
DELETE
FROM tasks
WHERE id = $1
  AND user_id = $2;

-- name: InsertTask :one
INSERT INTO tasks (content, user_id, created_date)
VALUES (@content, @user_id, @created_date)
RETURNING *;

-- name: CountTaskByUser :one
SELECT COUNT(id)
FROM tasks
WHERE user_id = $1
  AND created_date = CURRENT_DATE;
