package port

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/order/internal/application/core/entities"
)

type SeerService interface {
	GetSeerByUserUUID(ctx context.Context, userUUID uuid.UUID) (*entity.SeerResponse, error)
	GetUserUUIDBySeerUUID(ctx context.Context, seerUUID uuid.UUID) (*uuid.UUID, error)
}
