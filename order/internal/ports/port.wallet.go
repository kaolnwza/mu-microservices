package port

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/order/internal/application/core/entities"
)

type WalletService interface {
	GetUserWallet(ctx context.Context, userUUID uuid.UUID) (*entity.UserWallet, error)
	IncreaseUserWallet(ctx context.Context, userUUID uuid.UUID, fund int64) (*entity.UserWallet, error)
	DecreaseUserWallet(ctx context.Context, userUUID uuid.UUID, fund int64) (*entity.UserWallet, error)
}
