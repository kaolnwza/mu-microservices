package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/wallet/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/wallet/internal/ports"
)

type walSvc struct {
	tx   port.Transactor
	repo port.WalletRepository
}

func NewWalletService(tx port.Transactor, repo port.WalletRepository) port.WalletService {
	return &walSvc{tx: tx, repo: repo}
}

func (s *walSvc) CreateUserWallet(ctx context.Context, userUUID uuid.UUID) error {
	return s.repo.CreateWallet(ctx, userUUID)
}

func (s *walSvc) GetWalletByUserUUID(ctx context.Context, userUUID uuid.UUID) (*entity.UserWallet, error) {
	wallet := entity.UserWallet{}
	if err := s.repo.GetWalletByUserUUID(ctx, &wallet, userUUID); err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (s *walSvc) IncreaseWallet(ctx context.Context, userUUID uuid.UUID, amount int64) (*entity.UserWallet, error) {
	wallet := entity.UserWallet{}
	if err := s.repo.IncreaseWallet(ctx, &wallet, userUUID, amount); err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (s *walSvc) DecreaseWallet(ctx context.Context, userUUID uuid.UUID, amount int64) (*entity.UserWallet, error) {
	wallet := entity.UserWallet{}
	if err := s.repo.GetWalletByUserUUID(ctx, &wallet, userUUID); err != nil {
		return nil, err
	}

	fmt.Println(amount, wallet.Fund, int64(wallet.Fund), amount > wallet.Fund, amount > int64(wallet.Fund))

	if amount > int64(wallet.Fund) {
		return &wallet, errors.New("not enough fund")
	}

	if err := s.repo.DecreaseWallet(ctx, &wallet, userUUID, amount); err != nil {
		return nil, err
	}

	return &wallet, nil
}
