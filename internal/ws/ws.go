package ws

import (
	"Chat-Websocket/internal/db"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type IWsHandler interface {
	HandleWebSocket(c *gin.Context)
}

type IWsService interface {
	HandleConnection(room string, conn *websocket.Conn, token string) error
}

type IWsRepository interface {
	AddClient(ctx context.Context, room, username string, conn *websocket.Conn) error
	RemoveClient(room, username string, conn *websocket.Conn) error
	BroadcastMessage(msg Message, sender *websocket.Conn) error
	BroadcastToAll(msg Message) error
	GetClientCount(room string) int
	CheckRateLimit(conn *websocket.Conn) (bool, error)
	GetMessageHistory(room int32) []Message
	SaveMessage(msg Message) error
	GetUserByNameRepository(ctx context.Context, name string) (db.User, error)
	GetRoomByNameRepository(ctx context.Context, name string) (db.Room, error)
	GetUserByIDRepository(ctx context.Context, userID int32) (db.User, error)
}

type MessageInput struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

type Message struct {
	Room     string `json:"room"`
	Username string `json:"username"`
	Content  string `json:"content"`
}

const systemUser string = "SYSTEM"
