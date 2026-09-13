-- +goose Up
CREATE INDEX idx_user_username ON users (username);

-- +goose Down
DROP INDEX IF EXISTS idx_user_username;
