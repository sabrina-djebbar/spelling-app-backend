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

-- name: ListSetsByTags :many
SELECT  ss.id AS set_id, ss.name AS set_name, ss.description, ss.recommended_age, ss.tags as set_tags,ss.creator,
        sw.id AS word_id, sw.spelling, sw.definition, sw.difficulty, sw.total_available_points, sw.tags as word_tags, sw.class as word_class
FROM spelling_set ss JOIN spelling_set_words ssw ON ss.id = ssw.set_id JOIN spelling_word sw ON ssw.word_id = sw.id WHERE ss.tags LIKE '%' || $1 || '%';

-- name: ListWordsByTags :many
SELECT * FROM spelling_word WHERE tags LIKE  '%' || $1 || '%';

-- name: AddSpellingAttempt :one
INSERT INTO spelling_exercise (id, user_id, set_id, word_id,spelling, score, num_of_attempts,last_attempt) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING *;

-- name: GetSpellingExerciseAttemptByID :many
SELECT * FROM spelling_exercise se JOIN spelling_set ss ON se.set_id = ss.id JOIN spelling_word sw ON se.word_id = sw.id WHERE se.id = $1;

-- name: GetSpellingExerciseAttemptByUser :many
SELECT se.id as exercise_id,se.user_id, se.spelling as spelling_attempt, se.last_attempt, se.num_of_attempts, se.score, ss.id as set_id, ss.name AS set_name, ss.description, ss.recommended_age, ss.tags as set_tags,ss.creator,
       sw.id AS word_id, sw.spelling as correct_spelling, sw.definition, sw.difficulty, sw.total_available_points, sw.tags as word_tags, sw.class as word_class FROM spelling_exercise se JOIN spelling_set ss ON se.set_id = ss.id JOIN spelling_word sw ON se.word_id = sw.id WHERE se.user_id = $1;
