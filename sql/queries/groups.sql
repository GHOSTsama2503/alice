-- name: CreateGroup :exec
INSERT INTO groups (id, title, is_private)
VALUES (?, ?, ?);


-- name: GroupExists :many
SELECT id FROM groups
WHERE id = ?;
