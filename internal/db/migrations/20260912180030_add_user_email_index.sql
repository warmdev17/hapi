-- +goose Up
CREATE INDEX idx_user_email ON users (email);

-- +goose Down
DROP INDEX IF EXISTS idx_user_email;

