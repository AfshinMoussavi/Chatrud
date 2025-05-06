package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	chatService IChatService
	hub         *Hub
}

func NewHandler(h *Hub, s IChatService) *Handler {
	return &Handler{hub: h, chatService: s}
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room, err := h.chatService.CreateRoomService(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	h.hub.Rooms[strconv.Itoa(int(room.ID))] = room

	c.JSON(http.StatusOK, gin.H{"data": room})
}

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

	if roomID == "" || clientID == "" || username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "roomId, userId, or username is missing"})
		return
	}

	// Check if the room exists
	if _, ok := h.hub.Rooms[roomID]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "room not found"})
		return
	}

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
		Content:  fmt.Sprintf("%s has joined the room", username),
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

//func (h *Handler) Joinroom(c *gin.Context) {
//	// Upgrade HTTP connection to WebSocket
//	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	// Get room ID and client info from URL parameters
//	roomID := c.Param("roomId")
//	clientID := c.Query("userId")
//	username := c.Query("username")
//
//	// Create a new client instance
//	cl := &Client{
//		Conn:     conn,
//		Message:  make(chan *Message, 10),
//		ID:       clientID,
//		RoomID:   roomID,
//		Username: username,
//	}
//
//	// Notify others that a new user has joined
//	m := &Message{
//		Content:  "A new user has joined the room",
//		RoomID:   roomID,
//		Username: username,
//	}
//
//	// Register the new client and broadcast the join message
//	h.hub.Register <- cl
//	h.hub.Broadcast <- m
//
//	// Start goroutines for reading from and writing to the client
//	go cl.writeMessage()
//	cl.readMessage(h.hub)
//}

//
//// RoomRes represents a room for response purposes
//type RoomRes struct {
//	ID   string `json:"id"`
//	Name string `json:"name"`
//}
//
//// GetRooms returns a list of all rooms
//func (h *Handler) GetRooms(c *gin.Context) {
//	rooms := make([]RoomRes, 0)
//
//	for _, r := range h.hub.Rooms {
//		rooms = append(rooms, RoomRes{
//			ID:   r.ID,
//			Name: r.Name,
//		})
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": rooms})
//}
//
//// ClientRes represents a client for response purposes
//type ClientRes struct {
//	ID       string `json:"id"`
//	Username string `json:"username"`
//}
//
//// GetClients returns a list of all clients in a specific room
//func (h *Handler) GetClients(c *gin.Context) {
//	var clients []ClientRes
//	roomId := c.Param("roomId")
//
//	// If the room does not exist, return an empty list
//	if _, ok := h.hub.Rooms[roomId]; !ok {
//		clients = make([]ClientRes, 0)
//		c.JSON(http.StatusOK, gin.H{"message": clients})
//		return
//	}
//
//	// Collect all clients from the room
//	for _, c := range h.hub.Rooms[roomId].Clients {
//		clients = append(clients, ClientRes{
//			ID:       c.ID,
//			Username: c.Username,
//		})
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": clients})
//}
