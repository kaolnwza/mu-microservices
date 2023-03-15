package postgres

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/wallet/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/wallet/internal/ports"
)

type walRepo struct {
	tx port.Transactor
}

func NewWalletRepository(tx port.Transactor) port.WalletRepository {
	return &walRepo{tx: tx}
}

func (r *walRepo) CreateWallet(ctx context.Context, userUUID uuid.UUID) error {
	query := `
		INSERT INTO user_wallet (user_uuid)
		VALUES ($1)
	`

	return r.tx.Insert(ctx, query, userUUID)
}

func (r *walRepo) GetWalletByUserUUID(ctx context.Context, wallet *entity.UserWallet, userUUID uuid.UUID) error {
	query := `
		SELECT
			user_uuid,
			fund
		FROM user_wallet
		WHERE user_uuid = $1
	`

	return r.tx.GetOne(ctx, wallet, query, userUUID)
}

func (r *walRepo) IncreaseWallet(ctx context.Context, wallet *entity.UserWallet, userUUID uuid.UUID, amount int64) error {
	query := `
		UPDATE user_wallet
		SET 
			fund = fund + $2
		WHERE 
			user_uuid = $1
		RETURNING user_uuid, fund
	`

	return r.tx.Update(ctx, query, userUUID, amount)
}

func (r *walRepo) DecreaseWallet(ctx context.Context, wallet *entity.UserWallet, userUUID uuid.UUID, amount int64) error {
	query := `
		UPDATE user_wallet
		SET 
			fund = fund - $2
		WHERE 
			user_uuid = $1
		RETURNING user_uuid, fund
	`

	return r.tx.Update(ctx, query, userUUID, amount)
}
