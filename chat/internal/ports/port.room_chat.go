package port

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/chat/internal/application/core/entities"
)

type RoomChatRepository interface {
	NewMessage(ctx context.Context, userUUID uuid.UUID, roomUUID uuid.UUID, message string) error
	GetMessageByRoomUUID(ctx context.Context, dest *[]*entity.RoomChat, roomUUID uuid.UUID, offset int, limit int) error
}

type RoomChatService interface {
	NewMessage(ctx context.Context, userUUID uuid.UUID, roomUUID uuid.UUID, message string) error
	GetMessageByRoomUUID(ctx context.Context, roomUUID uuid.UUID, offset int, limit int) (*[]*entity.RoomChatResponse, error)
}
