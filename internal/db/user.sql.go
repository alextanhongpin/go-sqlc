// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
	name, age
) VALUES ($1, $2)
RETURNING id, name, age, created_at, updated_at, deleted_at, profile_id
`

type CreateUserParams struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Name, arg.Age)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Age,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.ProfileID,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, name, age, created_at, updated_at, deleted_at, profile_id
FROM users
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Age,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.ProfileID,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, name, age, created_at, updated_at, deleted_at, profile_id
FROM USERS
ORDER BY name
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Age,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.ProfileID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}