package port

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/chat/internal/application/core/entities"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, stUserUUID uuid.UUID, ndUserUUID uuid.UUID, startTime string, endTime string, orderUUID uuid.UUID) error
	GetRoomByUserUUID(ctx context.Context, dest *[]*entity.Room, userUUID uuid.UUID) error
	GetRoomByUUID(ctx context.Context, dest *entity.Room, userUUID uuid.UUID, roomUUID uuid.UUID) error
	UpdateRoomStatusByOrderUUID(ctx context.Context, orderUUID uuid.UUID, status bool) error
}

type RoomService interface {
	CreateRoom(ctx context.Context, stUserUUID uuid.UUID, ndUserUUID uuid.UUID, startTime string, endTime string, orderUUID uuid.UUID) error
	GetRoomByUserUUID(ctx context.Context, userUUID uuid.UUID) (*[]*entity.RoomResponse, error)
	GetRoomByUUID(ctx context.Context, userUUID uuid.UUID, roomUUID uuid.UUID) (*entity.RoomResponse, error)
	UpdateRoomStatusByOrderUUID(ctx context.Context, orderUUID uuid.UUID, status bool) error
}
