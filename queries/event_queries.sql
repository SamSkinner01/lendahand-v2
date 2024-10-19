-- name: GetAllFromEvents :many
SELECT * FROM events;

-- name: CreateEvent :one
INSERT INTO events (title, description, start_date, end_date, location) VALUES ($1, $2, $3, $4, $5) RETURNING *;