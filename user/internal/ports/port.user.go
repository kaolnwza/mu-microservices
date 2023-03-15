package port

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/user/internal/application/core/entities"
)

type UserRepository interface {
	GetByUUID(ctx context.Context, dest *entity.User, userUUID uuid.UUID) error
	CreateUser(ctx context.Context, dest *entity.User, display_name string, dob string, desc string) error
	Update(ctx context.Context, userUUID uuid.UUID, display_name string, dob string, desc string) error
}

type UserService interface {
	GetUserByUUID(ctx context.Context, userUUID uuid.UUID) (*entity.UserResponse, error)
	GetUserWithoutImgByUUID(ctx context.Context, userUUID uuid.UUID) (*entity.UserResponse, error)
	CreateUser(ctx context.Context, display_name string, dob string, desc string) (*entity.User, error)
	UpdateUser(ctx context.Context, userUUID uuid.UUID, display_name string, dob string, desc string, uploadUUID uuid.UUID) error
}
