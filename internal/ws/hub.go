package ws

// Room represents a chat room with connected clients
type Room struct {
	ID      string             `json:"id"`      // Unique identifier of the room
	Name    string             `json:"name"`    // Name of the room
	Clients map[string]*Client `json:"clients"` // Connected clients in the room
}

// Hub manages all rooms and handles client registration, unregistration, and broadcasting
type Hub struct {
	Rooms      map[string]*Room // All active rooms
	Register   chan *Client     // Channel for registering new clients
	Unregister chan *Client     // Channel for unregistering clients
	Broadcast  chan *Message    // Channel for broadcasting messages to clients
}

// NewHub creates and returns a new Hub instance
func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

// Run starts the hub's main event loop
func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			// Handle client registration
			if _, ok := h.Rooms[cl.RoomID]; ok {
				r := h.Rooms[cl.RoomID]

				if _, ok := r.Clients[cl.ID]; !ok {
					r.Clients[cl.ID] = cl
				}
			}

		case cl := <-h.Unregister:
			// Handle client unregistration
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					// If the room still has clients, broadcast user left message
					if len(h.Rooms[cl.RoomID].Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  "user left the chat",
							RoomID:   cl.RoomID,
							Username: cl.Username,
						}
					}
					// Remove client and close its message channel
					delete(h.Rooms[cl.RoomID].Clients, cl.ID)
					close(cl.Message)
				}
			}

		case m := <-h.Broadcast:
			// Handle broadcasting a message to all clients in a room
			if _, ok := h.Rooms[m.RoomID]; ok {
				for _, cl := range h.Rooms[m.RoomID].Clients {
					cl.Message <- m
				}
			}
		}
	}
}
