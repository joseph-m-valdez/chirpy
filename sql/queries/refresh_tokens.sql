-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens	(
	token, created_at, updated_at, user_id, expires_at, revoked_at
)	VALUES (
	$1,
	NOW(),
	NOW(),
	$2,
	$3,
	NULL
)
RETURNING token, user_id, created_at, updated_at, expires_at, revoked_at;

-- name: GetUserFromRefreshToken :one
SELECT u.id, u.created_at, u.updated_at, u.email, u.hashed_password
FROM refresh_tokens rt
JOIN users u ON u.id = rt.user_id
WHERE rt.token = $1
  AND rt.revoked_at IS NULL
  AND rt.expires_at > now();

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens rt
	SET revoked_at = NOW(), updated_at = NOW()
	WHERE rt.token = $1;
