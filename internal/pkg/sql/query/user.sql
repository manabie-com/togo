-- name: CreateUser :one
insert into users
(username, full_name, hashed_password, email, created_at, tier_id)
values ($1, $2, $3, $4, $5, $6)
returning *;

-- name: UpdateUserTier :one
update users
SET tier_id = $1
WHERE username = $2
returning *;

-- name: GetUserByName :one
select *
from users
where username = $1
limit 1;

-- name: ListUsers :many
select *
from users
order by created_at
limit $1 offset $2;

-- name: DeleteUser :exec
delete from users where username = $1;