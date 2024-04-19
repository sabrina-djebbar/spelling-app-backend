-- name: CreateUser :one
INSERT INTO users (username, date_of_birth, parent_code) VALUES ($1, $2, $3) RETURNING *;

-- name: CreateCredentials :exec
INSERT INTO credentials (user_id, password) VALUES($1,crypt($2,'crypt-des'));
