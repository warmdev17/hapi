-- +goose Up
CREATE TABLE IF NOT EXISTS couples (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
    user1_id uuid REFERENCES users (id),
    user2_id uuid REFERENCES users (id),
    start_date date,
    created_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user1_id, user2_id),
    CHECK (user1_id < user2_id),
    CHECK (user1_id <> user2_id)
);

-- +goose Down
DROP TABLE IF EXISTS couples;

