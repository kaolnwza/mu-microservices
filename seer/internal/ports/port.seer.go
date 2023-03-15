package port

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/seer/internal/application/core/entities"
)

type SeerRepository interface {
	GetByUUID(ctx context.Context, seer *entity.Seer, seerUUID uuid.UUID) error
	GetByUserUUID(ctx context.Context, seer *entity.Seer, userUUID uuid.UUID) error
}

type SeerService interface {
	GetSeerByUUID(ctx context.Context, seerUUID uuid.UUID) (*entity.SeerResponse, error)
	GetSeerByUserUUID(ctx context.Context, userUUID uuid.UUID) (*entity.SeerResponse, error)
}
