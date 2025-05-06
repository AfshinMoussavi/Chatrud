package ws

import "context"

type chatService struct {
	wsRepo IChatRepository
}

func NewService(repository IChatRepository) IChatService {
	return &chatService{wsRepo: repository}
}

func (s *chatService) CreateRoomService(ctx context.Context, req CreateRoomReq) (*Room, error) {
	roomRepo, err := s.wsRepo.CreateRoomRepository(ctx, req.Name)
	if err != nil {
		return &Room{}, err
	}

	room := &Room{
		ID:      roomRepo.ID,
		Name:    roomRepo.Name,
		Clients: make(map[string]*Client),
	}

	return room, nil
}
