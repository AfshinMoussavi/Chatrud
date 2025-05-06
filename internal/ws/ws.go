package ws

import (
	"Chat-Websocket/internal/db"
	"context"
	"github.com/gorilla/websocket"
	"net/http"
)

type IChatService interface {
	CreateRoomService(ctx context.Context, req CreateRoomReq) (*Room, error)
}

type IChatRepository interface {
	CreateRoomRepository(ctx context.Context, name string) (db.Room, error)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type CreateRoomReq struct {
	Name string `json:"name" binding:"required"`
}
