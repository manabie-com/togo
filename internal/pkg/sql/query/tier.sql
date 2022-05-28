-- name: CreateTier :one
insert into tiers
(id, name, description)
values ($1, $2, $3)
returning *;

-- name: GetTier :one
select *
from tiers
where id = $1
limit 1;

-- name: ListTiers :many
select *
from tiers
order by name
limit $1 offset $2;

-- name: DeleteTier :exec
delete from tiers where id = $1;