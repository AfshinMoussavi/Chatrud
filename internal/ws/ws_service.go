package ws

import (
	"Chat-Websocket/pkg/authPkg"
	"Chat-Websocket/pkg/loggerPkg"
	"context"
	"github.com/gorilla/websocket"
)

type wsService struct {
	repo   IWsRepository
	logger loggerPkg.ILogger
}

func NewService(repo IWsRepository, logger loggerPkg.ILogger) IWsService {
	return &wsService{repo: repo, logger: logger}
}

func (s *wsService) HandleConnection(room string, conn *websocket.Conn, token string) error {
	ctx := context.Background()

	// Validate token
	claims, err := authPkg.VerifyToken(token)
	if err != nil {
		s.logger.Error("Invalid token: %v", err)
		errMsg := Message{
			Room:     room,
			Username: systemUser,
			Content:  "Invalid or expired token",
		}
		if err := conn.WriteJSON(errMsg); err != nil {
			s.logger.Error("Error sending token error: %v", err)
		}
		conn.Close()
		return err
	}

	// Extract user ID from claims
	userID := claims.UserID
	if userID == 0 {
		s.logger.Error("Invalid user ID in token")
		errMsg := Message{
			Room:     room,
			Username: systemUser,
			Content:  "Invalid user ID in token",
		}
		if err := conn.WriteJSON(errMsg); err != nil {
			s.logger.Error("Error sending user ID error: %v", err)
		}
		conn.Close()
		return err
	}
	userID32 := int32(userID)
	// Get user from database using ID
	user, err := s.repo.GetUserByIDRepository(ctx, userID32)
	if err != nil {
		s.logger.Error("User not found: %v", err)
		errMsg := Message{
			Room:     room,
			Username: systemUser,
			Content:  "User not found",
		}
		if err := conn.WriteJSON(errMsg); err != nil {
			s.logger.Error("Error sending user not found error: %v", err)
		}
		conn.Close()
		return err
	}
	username := user.Name

	// Read initial message (optional, can be removed if username is from token)
	var initialMsg MessageInput
	if err := conn.ReadJSON(&initialMsg); err != nil {
		s.logger.Error("Error reading initial message in room %s: %v", room, err)
		errMsg := Message{
			Room:     room,
			Username: systemUser,
			Content:  "Failed to read initial message",
		}
		if err := conn.WriteJSON(errMsg); err != nil {
			s.logger.Error("Error sending initial message error: %v", err)
		}
		conn.Close()
		return err
	}

	// Validate initial username matches token username
	if initialMsg.Username != "" && initialMsg.Username != username {
		s.logger.Error("Initial username does not match token username")
		errMsg := Message{
			Room:     room,
			Username: systemUser,
			Content:  "Username does not match authenticated user",
		}
		if err := conn.WriteJSON(errMsg); err != nil {
			s.logger.Error("Error sending username mismatch error: %v", err)
		}
		conn.Close()
		return err
	}

	roomDB, err := s.repo.GetRoomByNameRepository(ctx, room)
	if err != nil {
		s.logger.Error("Room not found: %v", err)
		errMsg := Message{
			Room:     room,
			Username: systemUser,
			Content:  "Room not found",
		}
		if err := conn.WriteJSON(errMsg); err != nil {
			s.logger.Error("Error sending room not found error: %v", err)
		}
		conn.Close()
		return err
	}

	if err := s.repo.AddClient(ctx, room, username, conn); err != nil {
		s.logger.Error("Error adding client to room %s: %v", room, err)
		return err
	}

	history := s.repo.GetMessageHistory(roomDB.ID)
	for _, msg := range history {
		if err := conn.WriteJSON(msg); err != nil {
			s.logger.Error("Error sending history message in room %s: %v", room, err)
		}
	}

	joinMsg := Message{
		Room:     room,
		Username: systemUser,
		Content:  username + " joined to room",
	}

	if err := s.repo.BroadcastToAll(joinMsg); err != nil {
		s.logger.Error("Error broadcasting join message in room %s: %v", room, err)
	}
	if err := s.repo.SaveMessage(joinMsg); err != nil {
		s.logger.Error("Error saving join message in room %s: %v", room, err)
	}

	defer func() {
		leaveMsg := Message{
			Room:     room,
			Username: systemUser,
			Content:  username + " left from room",
		}
		if err := s.repo.BroadcastToAll(leaveMsg); err != nil {
			s.logger.Error("Error broadcasting leave message in room %s: %v", room, err)
		}
		if err := s.repo.SaveMessage(leaveMsg); err != nil {
			s.logger.Error("Error saving leave message in room %s: %v", room, err)
		}
		s.repo.RemoveClient(room, username, conn)
	}()

	for {
		var input MessageInput
		if err := conn.ReadJSON(&input); err != nil {
			s.logger.Error("Error reading message in room %s: %v", room, err)
			conn.Close()
			return err
		}

		if input.Content == "" {
			errMsg := Message{
				Room:     room,
				Username: systemUser,
				Content:  "Content is empty",
			}
			if err := conn.WriteJSON(errMsg); err != nil {
				s.logger.Error("Error sending content error: %v", err)
			}
			continue
		}

		if input.Username != username {
			errMsg := Message{
				Room:     room,
				Username: systemUser,
				Content:  "Username does not match authenticated user",
			}
			if err := conn.WriteJSON(errMsg); err != nil {
				s.logger.Error("Error sending username mismatch error: %v", err)
			}
			continue
		}

		allowed, err := s.repo.CheckRateLimit(conn)
		if err != nil {
			s.logger.Error("Error checking rate limit in room %s: %v", room, err)
			continue
		}
		if !allowed {
			errMsg := Message{
				Room:     room,
				Username: systemUser,
				Content:  "You have sent too many messages. Please wait a few seconds",
			}
			if err := conn.WriteJSON(errMsg); err != nil {
				s.logger.Error("Error sending rate limit error: %v", err)
			}
			continue
		}

		msg := Message{
			Room:     room,
			Username: username,
			Content:  input.Content,
		}

		if err := s.repo.BroadcastMessage(msg, conn); err != nil {
			s.logger.Error("Error broadcasting message in room %s: %v", room, err)
			continue
		}
		if err := s.repo.SaveMessage(msg); err != nil {
			s.logger.Error("Error saving message in room %s: %v", room, err)
			continue
		}
	}
	return nil
}
