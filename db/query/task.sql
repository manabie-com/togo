-- name: CreateTask :one
INSERT INTO
    tasks (owner, content, quantity)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetTask :one
SELECT
    *
FROM
    tasks
WHERE
    id = $1
LIMIT
    1;

-- name: ListTasksByOwner :many
SELECT
    *
FROM
    tasks
WHERE
    owner = $3
ORDER BY
    id
LIMIT
    $1
OFFSET
    $2;
