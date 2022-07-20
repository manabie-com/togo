-- name: GetTaskByID :one
SELECT
    *
FROM
    todo_tasks
WHERE
    id = ?;

-- name: GetTotalTaskByUserID :one
SELECT
    user_id,
    count(*) total_task
FROM
    todo_tasks
WHERE
    user_id = ?
    AND DATE(created_at) = DATE(NOW())
GROUP BY
    user_id;

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