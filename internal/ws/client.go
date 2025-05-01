package ws

import (
	"github.com/gorilla/websocket"
	"log"
)

// Client represents a single chatting user
type Client struct {
	Conn     *websocket.Conn // WebSocket connection for the client
	Message  chan *Message   // Channel to send messages to the client
	ID       string          `json:"id"`       // Unique client ID
	RoomID   string          `json:"roomId"`   // ID of the room the client belongs to
	Username string          `json:"username"` // Username of the client
}

// Message represents a chat message
type Message struct {
	Content  string `json:"content"`  // The message text
	RoomID   string `json:"roomId"`   // The room where the message was sent
	Username string `json:"username"` // The sender's username
}

// writeMessage listens for new messages on the Message channel
// and writes them to the WebSocket connection as JSON
func (c *Client) writeMessage() {
	defer func() {
		// Close the WebSocket connection when the function ends
		c.Conn.Close()
	}()

	for {
		// Wait for a new message to send
		message, ok := <-c.Message
		if !ok {
			// If the channel is closed, exit the loop
			return
		}
		// Send the message to the client as JSON
		c.Conn.WriteJSON(message)
	}
}

// readMessage continuously reads messages from the WebSocket connection
// and broadcasts them to all clients in the same room via the Hub
func (c *Client) readMessage(hub *Hub) {
	defer func() {
		// When exiting, unregister the client and close the connection
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		// Read a message from the WebSocket connection
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			// If the error is an unexpected close, log it
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			// Exit the loop if an error occurs
			break
		}

		// Create a new Message object from the received data
		msg := &Message{
			Content:  string(m),
			RoomID:   c.RoomID,
			Username: c.Username,
		}

		// Send the message to the hub for broadcasting
		hub.Broadcast <- msg
	}
}
