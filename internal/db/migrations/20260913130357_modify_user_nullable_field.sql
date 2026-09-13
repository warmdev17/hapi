-- +goose Up
ALTER TABLE users
    ALTER COLUMN hashed_password SET NOT NULL;

ALTER TABLE users
    ALTER COLUMN display_name SET NOT NULL;

-- +goose Down
ALTER TABLE users
    ALTER COLUMN hashed_password DROP NOT NULL;

ALTER TABLE users
    ALTER COLUMN display_name DROP NOT NULL;

