-- name: CreateUser :one
INSERT INTO users (username, email, hashed_password, display_name, birth_day, gender)
    VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    *;

