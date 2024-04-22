-- name: CreateUser :one
INSERT INTO users (id, username, date_of_birth, parent_code) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: CreateCredentials :exec
INSERT INTO credentials ( user_id, password) VALUES($1,crypt($2,'crypt-des'));

-- name: GetUser :one
SELECT * FROM users where id == $1;

-- name: ListUsers :many
SELECT * FROM users;