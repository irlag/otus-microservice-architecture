// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: users.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const CreateUser = `-- name: CreateUser :one
INSERT INTO users (
    email,
    first_name,
    last_name,
    password
) VALUES (
     $1, $2, $3, $4
 ) RETURNING id, email, first_name, last_name, created_at
`

type CreateUserParams struct {
	Email     string         `db:"email" json:"email"`
	FirstName string         `db:"first_name" json:"first_name"`
	LastName  sql.NullString `db:"last_name" json:"last_name"`
	Password  string         `db:"password" json:"password"`
}

type CreateUserRow struct {
	ID        int64          `db:"id" json:"id"`
	Email     string         `db:"email" json:"email"`
	FirstName string         `db:"first_name" json:"first_name"`
	LastName  sql.NullString `db:"last_name" json:"last_name"`
	CreatedAt time.Time      `db:"created_at" json:"created_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRowContext(ctx, CreateUser,
		arg.Email,
		arg.FirstName,
		arg.LastName,
		arg.Password,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
	)
	return i, err
}

const DeleteUser = `-- name: DeleteUser :execrows
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) (int64, error) {
	result, err := q.db.ExecContext(ctx, DeleteUser, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const GetUser = `-- name: GetUser :one
SELECT id, email, first_name, last_name, password, created_at, updated_at FROM users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, GetUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const GetUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, first_name, last_name, password, created_at, updated_at FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, GetUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const GetUserForUpdate = `-- name: GetUserForUpdate :one
SELECT id, email, first_name, last_name, password, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1
FOR UPDATE
`

func (q *Queries) GetUserForUpdate(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, GetUserForUpdate, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const UpdateUser = `-- name: UpdateUser :one
UPDATE users SET
    email = $2,
    first_name  = $3,
    last_name = $4,
    updated_at = now()
WHERE id = $1 RETURNING id, email, first_name, last_name, password, created_at, updated_at
`

type UpdateUserParams struct {
	ID        int64          `db:"id" json:"id"`
	Email     string         `db:"email" json:"email"`
	FirstName string         `db:"first_name" json:"first_name"`
	LastName  sql.NullString `db:"last_name" json:"last_name"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, UpdateUser,
		arg.ID,
		arg.Email,
		arg.FirstName,
		arg.LastName,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FirstName,
		&i.LastName,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
