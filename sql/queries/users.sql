-- name: CreateUser :one
INSERT INTO users (id, name, create_at, update_at)
VALUES ($1, $2, $3, $4)
RETURNING *;
