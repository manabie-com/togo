-- name: CreateTask :one
insert into tasks
(name, assignee, assign_date, description, status, creator, start_date, end_date)
values ($1, $2, $3, $4, $5, $6, $7, $8)
returning *;

-- name: AssignTask :one
update tasks
SET assignee = $2
WHERE id = $1
returning *;


-- name: GetTask :one
select *
from tasks
where id = $1
limit 1;

-- name: ListTasks :many
select *
from tasks
order by created_at
limit $1 offset $2;

-- name: CountTaskByAssigneeToday :one
SELECT count(assign_date)
FROM tasks
where assignee=$1 and DATE(assign_date) = current_date
group BY DATE(assign_date);

-- name: DeleteTask :exec
delete
from tasks
where id = $1;