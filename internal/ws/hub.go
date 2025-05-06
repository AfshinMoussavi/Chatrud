package ws

import "fmt"

type Room struct {
	ID      int32              `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			// Handle client registration
			if room, ok := h.Rooms[cl.RoomID]; ok {
				room.Clients[cl.ID] = cl
				h.Broadcast <- &Message{
					Content:  fmt.Sprintf("%s has joined the room", cl.Username),
					RoomID:   cl.RoomID,
					Username: cl.Username,
				}
			}

		case cl := <-h.Unregister:
			// Handle client unregistration
			if room, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := room.Clients[cl.ID]; ok {
					delete(room.Clients, cl.ID)
					close(cl.Message)
					h.Broadcast <- &Message{
						Content:  fmt.Sprintf("%s has left the room", cl.Username),
						RoomID:   cl.RoomID,
						Username: cl.Username,
					}
				}
			}

		case m := <-h.Broadcast:
			// Handle broadcasting a message to all clients in a room
			if room, ok := h.Rooms[m.RoomID]; ok {
				for _, cl := range room.Clients {
					select {
					case cl.Message <- m:
					default:
						close(cl.Message)
						delete(room.Clients, cl.ID)
					}
				}
			}
		}
	}
}
