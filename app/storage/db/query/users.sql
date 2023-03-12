-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserForUpdate :one
SELECT * FROM users
WHERE id = $1 LIMIT 1
FOR UPDATE;

-- name: CreateUser :one
INSERT INTO users (
    email,
    first_name,
    last_name,
    password
) VALUES (
     $1, $2, $3, $4
 ) RETURNING id, email, first_name, last_name, created_at;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users SET
    email = $2,
    first_name  = $3,
    last_name = $4,
    updated_at = now()
WHERE id = $1 RETURNING *;

-- name: DeleteUser :execrows
DELETE FROM users
WHERE id = $1;