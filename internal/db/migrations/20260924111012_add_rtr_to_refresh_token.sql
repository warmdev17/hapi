-- +goose Up
ALTER TABLE refresh_tokens RENAME COLUMN token TO token_hash;

ALTER TABLE refresh_tokens
    ADD COLUMN revoked_at timestamptz NULL;

UPDATE
    refresh_tokens
SET
    revoked_at = now()
WHERE
    is_revoked = TRUE;

ALTER TABLE refresh_tokens
    DROP COLUMN is_revoked;

ALTER TABLE refresh_tokens
    ADD COLUMN family_id uuid NOT NULL DEFAULT gen_random_uuid (),
    ADD COLUMN is_used boolean DEFAULT FALSE,
    ADD COLUMN parent_id uuid REFERENCES refresh_tokens (id) ON DELETE SET NULL,
    ADD COLUMN user_agent text NULL,
    ADD COLUMN ip_address inet NULL;

ALTER TABLE refresh_tokens
    ALTER COLUMN family_id DROP DEFAULT;

CREATE INDEX idx_refresh_tokens_user_family ON refresh_tokens (user_id, family_id);

-- +goose Down
DROP INDEX IF EXISTS idx_refresh_tokens_user_family;

ALTER TABLE refresh_tokens
    DROP COLUMN IF EXISTS ip_address,
    DROP COLUMN IF EXISTS user_agent,
    DROP COLUMN IF EXISTS parent_id,
    DROP COLUMN IF EXISTS is_used,
    DROP COLUMN IF EXISTS family_id;

ALTER TABLE refresh_tokens
    ADD COLUMN is_revoked boolean NOT NULL DEFAULT FALSE;

UPDATE
    refresh_tokens
SET
    is_revoked = TRUE
WHERE
    revoked_at IS NOT NULL;

ALTER TABLE refresh_tokens
    DROP COLUMN IF EXISTS revoked_at;

ALTER TABLE refresh_tokens RENAME COLUMN token_hash TO token;

