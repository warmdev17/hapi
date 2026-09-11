-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (family_id, user_id, token_hash)
    VALUES ($1, $2, $3)
RETURNING
    *;

-- name: RevokeToken :exec
UPDATE
    refresh_tokens
SET
    is_revoked = TRUE
WHERE
    id = $1
    AND token_hash = $2;

