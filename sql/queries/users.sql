-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (
	gen_random_uuid(),
	NOW(),
	NOW(),
	$1,
	$2
)
RETURNING *;

-- name: GetUser :one
SELECT id, created_at, updated_at, email, hashed_password, is_chirpy_red
	FROM users WHERE email = $1;

-- name: UpdateUserAuth :one
UPDATE users
SET
	email = $1,
	hashed_password = $2,
	updated_at = NOW()
WHERE id = $3
RETURNING id, created_at, updated_at, email, hashed_password, is_chirpy_red;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: UpdateUserMembership :one
UPDATE users
SET
	is_chirpy_red = $1,
	updated_at = NOW()
WHERE id = $2
RETURNING id, created_at, updated_at, email, is_chirpy_red;

