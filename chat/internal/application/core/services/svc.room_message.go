package service

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/chat/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/chat/internal/ports"
	"github.com/kaolnwza/muniverse/chat/lib/helper"
)

type roomMsgSvc struct {
	tx          port.Transactor
	roomMsgRepo port.RoomMessageRepository
}

func NewRoomMessageService(tx port.Transactor, roomMsgRepo port.RoomMessageRepository) port.RoomMessageService {
	return &roomMsgSvc{
		tx:          tx,
		roomMsgRepo: roomMsgRepo,
	}
}

func (s *roomMsgSvc) CreateRoomMessage(ctx context.Context, roomUUID uuid.UUID, userUUID uuid.UUID, message string) error {
	return s.roomMsgRepo.CreateRoomMessage(ctx, roomUUID, userUUID, message)
}

func (s *roomMsgSvc) GetMessageByRoomUUID(ctx context.Context, roomUUID uuid.UUID, timeOffset string) (*[]*entity.RoomMessageResponse, error) {
	room := make([]*entity.RoomMessage, 0)
	if err := s.roomMsgRepo.GetMessageByRoomUUID(ctx, &room, roomUUID, timeOffset); err != nil {
		return nil, err
	}

	resp := make([]*entity.RoomMessageResponse, 0)
	if err := helper.StructCopy(room, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
