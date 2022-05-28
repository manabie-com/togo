-- name: CreateLimit :one
insert into limits
(tier_id, action, value)
values ($1, $2, $3)
returning *;

-- name: GetLimit :one
select *
from limits
where tier_id = $1 and action = $2
limit 1;

-- name: ListLimits :many
select *
from limits
limit $1 offset $2;

-- name: DeleteLimit :exec
delete from limits where tier_id = $1 and action = $2;