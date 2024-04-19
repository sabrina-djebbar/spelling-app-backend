-- name: CreateUser :one
INSERT INTO users (id, username, date_of_birth, parent_code) VALUES ($1, $2, $3, $4) RETURNING id;

-- name: CreateCredentials :exec
INSERT INTO credentials (id, user_id, password) VALUES($1,$2,crypt($3,'crypt-des'));
