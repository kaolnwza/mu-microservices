package port

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/wallet/internal/application/core/entities"
)

type WalletRepository interface {
	CreateWallet(ctx context.Context, userUUID uuid.UUID) error
	GetWalletByUserUUID(ctx context.Context, wallet *entity.UserWallet, userUUID uuid.UUID) error
	IncreaseWallet(ctx context.Context, wallet *entity.UserWallet, userUUID uuid.UUID, amount int64) error
	DecreaseWallet(ctx context.Context, wallet *entity.UserWallet, userUUID uuid.UUID, amount int64) error
}

type WalletService interface {
	CreateUserWallet(ctx context.Context, userUUID uuid.UUID) error
	GetWalletByUserUUID(ctx context.Context, userUUID uuid.UUID) (*entity.UserWallet, error)
	IncreaseWallet(ctx context.Context, userUUID uuid.UUID, amount int64) (*entity.UserWallet, error)
	DecreaseWallet(ctx context.Context, userUUID uuid.UUID, amount int64) (*entity.UserWallet, error)
}
