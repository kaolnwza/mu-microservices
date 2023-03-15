package port

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/seer/internal/application/core/entities"
)

type UserService interface {
	GetUserInfo(ctx context.Context, userUUID uuid.UUID) (*entity.UserResponse, error)
}
