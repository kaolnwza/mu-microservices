package port

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/chat/internal/application/core/entities"
)

type RoomMessageRepository interface {
	CreateRoomMessage(ctx context.Context, roomUUID uuid.UUID, userUUID uuid.UUID, message string) error
	GetMessageByRoomUUID(ctx context.Context, msg *[]*entity.RoomMessage, roomUUID uuid.UUID, timeOffset string) error
}

type RoomMessageService interface {
	CreateRoomMessage(ctx context.Context, roomUUID uuid.UUID, userUUID uuid.UUID, message string) error
	GetMessageByRoomUUID(ctx context.Context, roomUUID uuid.UUID, timeOffset string) (*[]*entity.RoomMessageResponse, error)
}
