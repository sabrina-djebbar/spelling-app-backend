-- name: CreateUser :one
INSERT INTO users (id, username, date_of_birth, parent_code) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: CreateCredentials :exec
INSERT INTO credentials (id, user_id, password) VALUES($1,$2,crypt($3,'crypt-des'));

-- name: GetUser :one
SELECT * FROM users where id = $1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: FindByUsername :one
SELECT * FROM users WHERE username = $1;

-- name: FindCredentials :one
SELECT id FROM credentials WHERE user_id = $1 and password = crypt($2, 'crypt-des');

-- name: DeleteUser :exec
DELETE FROM users WHERE id == $1;

-- name: UpdateParentCode :one
UPDATE users
SET parent_code = $2
WHERE id = $1
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET date_of_birth = $2, username = $3
WHERE id = $1
RETURNING *;
