package ws

import (
	"Chat-Websocket/internal/db"
	"context"
)

type chatRepository struct {
	query *db.Queries
}

func NewRepository(query *db.Queries) IChatRepository {
	return &chatRepository{query: query}
}

func (r *chatRepository) CreateRoomRepository(ctx context.Context, name string) (db.Room, error) {
	room, err := r.query.CreateRoom(ctx, name)
	if err != nil {
		return db.Room{}, err
	}
	return room, nil
}
