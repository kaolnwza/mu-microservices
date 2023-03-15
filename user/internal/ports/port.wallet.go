package port

import (
	"context"

	"github.com/google/uuid"
)

type WalletService interface {
	CreateNewUserWallet(ctx context.Context, userUUID uuid.UUID) error
}
