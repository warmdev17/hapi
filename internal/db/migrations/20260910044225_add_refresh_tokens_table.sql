-- +goose Up
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id uuid REFERENCES users (id) NOT NULL,
    token_hash char(64) NOT NULL UNIQUE,
    family_id uuid NOT NULL,
    is_revoked boolean NOT NULL DEFAULT FALSE,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP + interval '7 days'
);

CREATE INDEX idx_valid_tokens_by_user ON refresh_tokens (user_id, is_revoked, expires_at);

-- +goose Down
DROP TABLE IF EXISTS refresh_token;

DROP INDEX IF EXISTS idx_valid_tokens_by_user;

