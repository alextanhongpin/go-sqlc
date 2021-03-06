// Code generated by sqlc. DO NOT EDIT.
// source: profile.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const getUserProfile = `-- name: GetUserProfile :one
SELECT users.id, users.name, users.age, users.created_at, users.updated_at, users.deleted_at, users.profile_id, profiles.id, profiles.name, profiles.created_at, profiles.updated_at, profiles.deleted_at
FROM users
JOIN profiles ON (users.id = profiles.user_id)
WHERE users.deleted_at IS NULL
	and profiles.deleted_at IS NULL
	and users.id = $1 LIMIT 1
`

type GetUserProfileRow struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Age         int32        `json:"age"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
	ProfileID   uuid.UUID    `json:"profile_id"`
	ID_2        uuid.UUID    `json:"id_2"`
	Name_2      string       `json:"name_2"`
	CreatedAt_2 time.Time    `json:"created_at_2"`
	UpdatedAt_2 time.Time    `json:"updated_at_2"`
	DeletedAt_2 sql.NullTime `json:"deleted_at_2"`
}

func (q *Queries) GetUserProfile(ctx context.Context, id uuid.UUID) (GetUserProfileRow, error) {
	row := q.db.QueryRowContext(ctx, getUserProfile, id)
	var i GetUserProfileRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Age,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.ProfileID,
		&i.ID_2,
		&i.Name_2,
		&i.CreatedAt_2,
		&i.UpdatedAt_2,
		&i.DeletedAt_2,
	)
	return i, err
}
