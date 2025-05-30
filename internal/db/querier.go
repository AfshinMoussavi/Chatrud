// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"
)

type Querier interface {
	CreateChat(ctx context.Context, arg CreateChatParams) (Chat, error)
	CreateRoom(ctx context.Context, name string) (Room, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteChat(ctx context.Context, id int32) error
	DeleteRoom(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id int32) error
	GetChatByID(ctx context.Context, id int32) (Chat, error)
	GetChatsByRoom(ctx context.Context, roomID int32) ([]Chat, error)
	GetChatsByUserAndRoom(ctx context.Context, arg GetChatsByUserAndRoomParams) ([]Chat, error)
	GetChatsByUserID(ctx context.Context, senderID int32) ([]Chat, error)
	GetRoomByID(ctx context.Context, id int32) (Room, error)
	GetRoomByName(ctx context.Context, name string) (Room, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id int32) (User, error)
	GetUserByName(ctx context.Context, name string) (User, error)
	GetUserByPhone(ctx context.Context, phone string) (User, error)
	ListRooms(ctx context.Context) ([]Room, error)
	ListUsers(ctx context.Context) ([]User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
}

var _ Querier = (*Queries)(nil)
