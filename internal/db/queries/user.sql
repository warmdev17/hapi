-- name: CreateUser :one
INSERT INTO users (username, email, hashed_password, display_name, birth_day, gender)
    VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    *;

-- name: GetUserByEmail :one
SELECT
    *
FROM
    users
WHERE
    email = $1;

-- name: CheckEmailExists :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            users
        WHERE
            email = $1);

