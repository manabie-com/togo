-- name: CreateTask :one
INSERT INTO tasks (
  content,
  user_id,
  created_date
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: RetrieveTasks :many
SELECT * FROM tasks
WHERE user_id = $1 
AND created_date = $2
ORDER BY created_date;