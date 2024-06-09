-- name: CreateUser :one
INSERT INTO users (id, language_code)
VALUES (?, ?)
RETURNING *;


-- name: GetUserLocale :one
SELECT language_code FROM users
WHERE id = ?;


-- name: SetUserLocale :exec
UPDATE users SET language_code = ?
WHERE id = ?;
