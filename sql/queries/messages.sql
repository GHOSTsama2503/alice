-- name: CreateMessages :exec
INSERT INTO messages (group_id, welcome_message)
VALUES (?, ?);


-- name: GetWelcomeMessage :one
SELECT welcome_message FROM messages
WHERE group_id = ?;


-- name: SetWelcomeMessage :exec
UPDATE messages
SET welcome_message = ?
WHERE group_id = ?;
