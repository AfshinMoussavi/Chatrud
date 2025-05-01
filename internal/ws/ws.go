package ws

import "github.com/gin-gonic/gin"

type HubInterface interface {
	Run()
	RegisterClient(c *Client)
	UnregisterClient(c *Client)
	BroadcastMessage(m *Message)
	GetRooms() []*Room
	GetClients(roomID string) []*Client
	CreateRoom(room *Room)
}

type ClientInterface interface {
	ReadMessage(hub HubInterface)
	WriteMessage()
}

type HandlerInterface interface {
	CreateRoom(c *gin.Context)
	JoinRoom(c *gin.Context)
	GetRooms(c *gin.Context)
	GetClients(c *gin.Context)
}
