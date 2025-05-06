-- name: CreateUser :one
INSERT INTO users (name, email, phone, password)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetUserByPhone :one
SELECT * FROM users
WHERE phone = $1;

-- name: GetUserByName :one
SELECT * FROM users
WHERE name = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: UpdateUser :one
UPDATE users
SET name = $2,email = $3,phone = $4,updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: UpdateUserPassword :exec
UPDATE users
SET password = $2,updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: CreateRoom :one
INSERT INTO rooms (name)
VALUES ($1)
RETURNING *;

-- name: GetRoomByID :one
SELECT * FROM rooms
WHERE id = $1;

-- name: ListRooms :many
SELECT * FROM rooms
ORDER BY id;

-- name: DeleteRoom :exec
DELETE FROM rooms
WHERE id = $1;

-- name: CreateChat :one
INSERT INTO chats (room_id, sender_id, message)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetChatByID :one
SELECT * FROM chats
WHERE id = $1;

-- name: ListChatsByRoom :many
SELECT * FROM chats
WHERE room_id = $1
ORDER BY created_at;

-- name: DeleteChat :exec
DELETE FROM chats
WHERE id = $1;

-- name: GetChatsByUserID :many
SELECT * FROM chats
WHERE sender_id = $1
ORDER BY created_at;

-- name: GetChatsByUserAndRoom :many
SELECT * FROM chats
WHERE sender_id = $1 AND room_id = $2
ORDER BY created_at;
