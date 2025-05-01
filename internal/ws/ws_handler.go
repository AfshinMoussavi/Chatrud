package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// Handler handles HTTP requests related to the WebSocket Hub
type Handler struct {
	hub *Hub
}

// NewHandler creates a new Handler with the given Hub
func NewHandler(h *Hub) *Handler {
	return &Handler{hub: h}
}

// CreateRoomReq represents the expected request body for creating a room
type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// CreateRoom handles the creation of a new chat room
func (h *Handler) CreateRoom(c *gin.Context) {
	var req CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		// If request body is invalid, return a 400 error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new room and add it to the hub
	h.hub.Rooms[req.ID] = &Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: map[string]*Client{},
	}

	c.JSON(http.StatusOK, gin.H{"message": req})
}

// upgrader upgrades HTTP connections to WebSocket connections
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow any origin (for testing purposes)
		return true
	},
}

// JoinRoom upgrades the connection to WebSocket and registers a new client in a room
func (h *Handler) JoinRoom(c *gin.Context) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get room ID and client info from URL parameters
	roomID := c.Param("roomId")
	clientID := c.Query("userId")
	username := c.Query("username")

	// Create a new client instance
	cl := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	// Notify others that a new user has joined
	m := &Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: username,
	}

	// Register the new client and broadcast the join message
	h.hub.Register <- cl
	h.hub.Broadcast <- m

	// Start goroutines for reading from and writing to the client
	go cl.writeMessage()
	cl.readMessage(h.hub)
}

// RoomRes represents a room for response purposes
type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetRooms returns a list of all rooms
func (h *Handler) GetRooms(c *gin.Context) {
	rooms := make([]RoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": rooms})
}

// ClientRes represents a client for response purposes
type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

// GetClients returns a list of all clients in a specific room
func (h *Handler) GetClients(c *gin.Context) {
	var clients []ClientRes
	roomId := c.Param("roomId")

	// If the room does not exist, return an empty list
	if _, ok := h.hub.Rooms[roomId]; !ok {
		clients = make([]ClientRes, 0)
		c.JSON(http.StatusOK, gin.H{"message": clients})
		return
	}

	// Collect all clients from the room
	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientRes{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": clients})
}
