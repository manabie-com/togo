-- name: GetTaskByID :one
SELECT
    *
FROM
    todo_tasks
WHERE
    id = ?;

-- name: InsertTask :execresult
INSERT INTO
    todo_tasks(user_id, task_name)
VALUES
    (?, ?);