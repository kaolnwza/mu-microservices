package port

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ChatService interface {
	CreateChatRoom(ctx context.Context, orderUUID uuid.UUID, userUUID uuid.UUID, seerUUID uuid.UUID, startTime time.Time, endTime time.Time) error
}
