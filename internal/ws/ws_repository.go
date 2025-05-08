package ws

import (
	"Chat-Websocket/internal/db"
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

type wsRepository struct {
	queries     *db.Queries
	conns       map[*websocket.Conn]struct{ username, room string } // Track username and room for each connection
	usernames   map[string]*websocket.Conn                          // Track connection for each username
	rateLimits  map[*websocket.Conn][]time.Time                     // Map for rate limiting
	mutex       sync.RWMutex
	maxMessages int
	windowSize  time.Duration
}

func NewRepository(queries *db.Queries) IWsRepository {
	return &wsRepository{
		queries:     queries,
		conns:       make(map[*websocket.Conn]struct{ username, room string }),
		usernames:   make(map[string]*websocket.Conn),
		rateLimits:  make(map[*websocket.Conn][]time.Time),
		maxMessages: 3,               // Max 10 messages for rate limiting
		windowSize:  5 * time.Second, // In 5 seconds
	}
}

func (w *wsRepository) AddClient(ctx context.Context, room, username string, conn *websocket.Conn) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	// Check if username is already in use
	if existingConn, exists := w.usernames[username]; exists && existingConn != conn {
		// Close existing connection
		existingConn.Close()
		delete(w.conns, existingConn)
		delete(w.rateLimits, existingConn)
	}

	// Store connection info with room
	w.conns[conn] = struct{ username, room string }{username, room}
	w.usernames[username] = conn
	w.rateLimits[conn] = []time.Time{}

	return nil
}

func (w *wsRepository) RemoveClient(room, username string, conn *websocket.Conn) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	// Remove connection info
	delete(w.conns, conn)
	if w.usernames[username] == conn {
		delete(w.usernames, username)
	}
	delete(w.rateLimits, conn)

	return nil
}

// Exept sender
func (w *wsRepository) BroadcastMessage(msg Message, sender *websocket.Conn) error {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	// Broadcast to all clients in the room, except sender
	for conn, info := range w.conns {
		if conn != sender && info.room == msg.Room {
			if err := conn.WriteJSON(msg); err != nil {
				return err
			}
		}
	}
	return nil
}

func (w *wsRepository) BroadcastToAll(msg Message) error {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	// Broadcast to all clients in the room
	for conn, info := range w.conns {
		if info.room == msg.Room {
			if err := conn.WriteJSON(msg); err != nil {
				return err
			}
		}
	}
	return nil
}

func (w *wsRepository) GetClientCount(room string) int {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	count := 0
	for _, info := range w.conns {
		if info.room == room {
			count++
		}
	}
	return count
}

func (w *wsRepository) CheckRateLimit(conn *websocket.Conn) (bool, error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	_, exists := w.conns[conn]
	if !exists {
		return false, nil
	}

	now := time.Now()
	timestamps, exists := w.rateLimits[conn]
	if !exists {
		w.rateLimits[conn] = []time.Time{now}
		return true, nil
	}

	// Remove timestamps older than windowSize
	newTimestamps := []time.Time{}
	for _, t := range timestamps {
		if now.Sub(t) <= w.windowSize {
			newTimestamps = append(newTimestamps, t)
		}
	}

	// Add new timestamp
	newTimestamps = append(newTimestamps, now)
	w.rateLimits[conn] = newTimestamps

	// Check if within limit
	if len(newTimestamps) > w.maxMessages {
		return false, nil
	}
	return true, nil
}

func (w *wsRepository) GetMessageHistory(roomID int32) []Message {
	ctx := context.Background()
	dbMessages, err := w.queries.GetChatsByRoom(ctx, roomID)
	if err != nil {
		return []Message{}
	}

	messages := make([]Message, 0, len(dbMessages))
	for _, m := range dbMessages {
		// Get username from users table using SenderID
		user, err := w.queries.GetUserByID(ctx, m.SenderID)
		if err != nil {
			continue // Skip if user not found
		}

		if user.Name == "SYSTEM" {
			continue
		}

		messages = append(messages, Message{
			Room:     fmt.Sprintf("%d", m.RoomID),
			Username: user.Name,
			Content:  m.Message,
		})
	}
	return messages
}

func (w *wsRepository) SaveMessage(msg Message) error {
	ctx := context.Background()

	// Get room
	roomDB, err := w.queries.GetRoomByName(ctx, msg.Room)
	if err != nil {
		return err
	}

	// Get user
	userDB, err := w.queries.GetUserByName(ctx, msg.Username)
	if err != nil {
		return err
	}

	// Save message
	_, err = w.queries.CreateChat(ctx, db.CreateChatParams{
		RoomID:   roomDB.ID,
		SenderID: userDB.ID,
		Message:  msg.Content,
	})
	if err != nil {
		return err
	}
	return nil
}

func (w *wsRepository) GetUserByNameRepository(ctx context.Context, name string) (db.User, error) {
	user, err := w.queries.GetUserByName(ctx, name)
	if err != nil {
		return db.User{}, errors.New("user get by name failed")
	}
	return user, nil
}

func (w *wsRepository) GetRoomByNameRepository(ctx context.Context, name string) (db.Room, error) {
	room, err := w.queries.GetRoomByName(ctx, name)
	if err != nil {
		return db.Room{}, errors.New("room get by name failed")
	}
	return room, nil
}

func (w *wsRepository) GetUserByIDRepository(ctx context.Context, userID int32) (db.User, error) {
	user, err := w.queries.GetUserByID(ctx, userID)
	if err != nil {
		return db.User{}, errors.New("user get by ID failed")
	}
	return user, nil
}
