// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: user.sql

package db

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
insert into users
(username, full_name, hashed_password, email, created_at, tier_id)
values ($1, $2, $3, $4, $5, $6)
returning username, full_name, hashed_password, email, created_at, tier_id
`

type CreateUserParams struct {
	Username       string    `json:"username"`
	FullName       string    `json:"full_name"`
	HashedPassword string    `json:"hashed_password"`
	Email          string    `json:"email"`
	CreatedAt      time.Time `json:"created_at"`
	TierID         int32     `json:"tier_id"`
}

func (q *Queries) CreateUser(ctx context.Context, arg *CreateUserParams) (*User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.Username,
		arg.FullName,
		arg.HashedPassword,
		arg.Email,
		arg.CreatedAt,
		arg.TierID,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.FullName,
		&i.HashedPassword,
		&i.Email,
		&i.CreatedAt,
		&i.TierID,
	)
	return &i, err
}

const deleteUser = `-- name: DeleteUser :exec
delete from users where username = $1
`

func (q *Queries) DeleteUser(ctx context.Context, username string) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, username)
	return err
}

const getUserByName = `-- name: GetUserByName :one
select username, full_name, hashed_password, email, created_at, tier_id
from users
where username = $1
limit 1
`

func (q *Queries) GetUserByName(ctx context.Context, username string) (*User, error) {
	row := q.queryRow(ctx, q.getUserByNameStmt, getUserByName, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.FullName,
		&i.HashedPassword,
		&i.Email,
		&i.CreatedAt,
		&i.TierID,
	)
	return &i, err
}

const listUsers = `-- name: ListUsers :many
select username, full_name, hashed_password, email, created_at, tier_id
from users
order by created_at
limit $1 offset $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg *ListUsersParams) ([]*User, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.Username,
			&i.FullName,
			&i.HashedPassword,
			&i.Email,
			&i.CreatedAt,
			&i.TierID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserTier = `-- name: UpdateUserTier :one
update users
SET tier_id = $1
WHERE username = $2
returning username, full_name, hashed_password, email, created_at, tier_id
`

type UpdateUserTierParams struct {
	TierID   int32  `json:"tier_id"`
	Username string `json:"username"`
}

func (q *Queries) UpdateUserTier(ctx context.Context, arg *UpdateUserTierParams) (*User, error) {
	row := q.queryRow(ctx, q.updateUserTierStmt, updateUserTier, arg.TierID, arg.Username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.FullName,
		&i.HashedPassword,
		&i.Email,
		&i.CreatedAt,
		&i.TierID,
	)
	return &i, err
}
