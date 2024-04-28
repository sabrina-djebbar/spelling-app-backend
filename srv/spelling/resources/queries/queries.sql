-- name: CreateSpellingWord :one
INSERT INTO spelling_word (id, spelling, difficulty, definition, total_available_points, class, tags) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: CreateSpellingSet :one
INSERT INTO spelling_set(id, name, recommended_age, description, tags, creator) VALUES ($1, $2,$3,$4,$5, $6) RETURNING *;

-- name: AddWordToSet :exec
INSERT INTO spelling_set_words(set_id, word_id) VALUES ($1, $2);

-- name: GetWordDifficulty :one
SELECT difficulty from spelling_word where id = $1;

-- name: GetSpellingWord :one
SELECT * FROM spelling_word where id = $1;

