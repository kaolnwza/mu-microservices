package service

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/chat/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/chat/internal/ports"
	"github.com/kaolnwza/muniverse/chat/lib/helper"
)

type roomSvc struct {
	tx       port.Transactor
	roomRepo port.RoomRepository
}

func NewRoomService(tx port.Transactor, roomRepo port.RoomRepository) port.RoomService {
	return &roomSvc{tx: tx, roomRepo: roomRepo}
}

func (s *roomSvc) CreateRoom(ctx context.Context, stUserUUID uuid.UUID, ndUserUUID uuid.UUID, startTime string, endTime string) error {
	if err := s.roomRepo.CreateRoom(ctx, stUserUUID, ndUserUUID, startTime, endTime); err != nil {
		return err
	}

	return nil
}

func (s *roomSvc) GetRoomByUserUUID(ctx context.Context, userUUID uuid.UUID) (*[]*entity.RoomResponse, error) {
	rooms := make([]*entity.Room, 0)
	if err := s.roomRepo.GetRoomByUserUUID(ctx, &rooms, userUUID); err != nil {
		return nil, err
	}

	resp := make([]*entity.RoomResponse, 0)
	if err := helper.StructCopy(rooms, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *roomSvc) GetRoomByUUID(ctx context.Context, userUUID uuid.UUID, roomUUID uuid.UUID) (*entity.RoomResponse, error) {
	room := entity.Room{}
	if err := s.roomRepo.GetRoomByUUID(ctx, &room, userUUID, roomUUID); err != nil {
		return nil, err
	}

	resp := entity.RoomResponse{}
	if err := helper.StructCopy(room, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *roomSvc) UpdateRoomStatusByOrderUUID(ctx context.Context, orderUUID uuid.UUID, status bool) error {
	return s.roomRepo.UpdateRoomStatusByOrderUUID(ctx, orderUUID, status)
}
