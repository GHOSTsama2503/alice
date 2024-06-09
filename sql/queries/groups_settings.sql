-- name: CreateGroupSettings :exec
INSERT INTO groups_settings (
	group_id,
	is_welcome_message_enabled
)
VALUES (?, ?);


-- name: IsWelcomeMessageEnabled :one
SELECT is_welcome_message_enabled FROM groups_settings
WHERE group_id = ?;


-- name: EnableWelcomeMessage :exec
UPDATE groups_settings
SET is_welcome_message_enabled = 1
WHERE group_id = ?;


-- name: DisableWelcomeMessage :exec
UPDATE groups_settings
SET is_welcome_message_enabled = 0
WHERE group_id = ?;
