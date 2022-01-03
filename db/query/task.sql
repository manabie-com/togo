-- name: CreateTask :one
INSERT INTO
	tasks (NAME, OWNER, CONTENT)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetTask :one
SELECT
	*
FROM
	tasks
WHERE
	id = $1 AND
	deleted = FALSE
LIMIT
	1;

-- name: GetTaskByName :one
SELECT
	*
FROM
	tasks
WHERE
	NAME = $1 AND
	OWNER = $2 AND
	deleted = FALSE
LIMIT
	1;

-- name: ListTasksByOwner :many
SELECT
	*
FROM
	tasks
WHERE
	OWNER = $3 AND
	deleted = FALSE
ORDER BY
	id
LIMIT
	$1
OFFSET
	$2;

-- name: UpdateTaskByName :one
UPDATE
	tasks
SET
	CONTENT = $2,
	content_change_at = NOW()
WHERE
	NAME = $1 AND
	OWNER = $3 RETURNING *;

-- name: DeleteTaskByName :exec
UPDATE
	tasks
SET
	deleted = TRUE,
	deleted_at = NOW()
WHERE
	NAME = $1 AND
	OWNER = $2;

-- name: CountTasksCreatedToday :one
SELECT
	COUNT(content_change_at)
FROM
	tasks
WHERE
	OWNER = $1 AND
	content_change_at :: DATE = NOW() :: DATE;
