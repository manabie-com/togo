-- name: GetTaskByID :one
SELECT
    *
FROM
    todo_tasks
WHERE
    id = ?;

-- name: GetTaskByUserID :many
SELECT
    *
FROM
    todo_tasks
WHERE
    user_id = ?;

-- name: InsertTask :execresult
INSERT INTO
    todo_tasks(user_id, task_name)
VALUES
    (?, ?);