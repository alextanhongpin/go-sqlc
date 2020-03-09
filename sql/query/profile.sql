-- name: GetUserProfile :one
SELECT users.*, profiles.*
FROM users
JOIN profiles ON (users.id = profiles.user_id)
WHERE users.deleted_at IS NULL
	and profiles.deleted_at IS NULL
	and users.id = $1 LIMIT 1;
