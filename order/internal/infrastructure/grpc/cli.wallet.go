package rpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/kaolnwza/muniverse/order/internal/adapters/grpc/proto/pb"
	entity "github.com/kaolnwza/muniverse/order/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/order/internal/ports"
)

type walletRpcCli struct {
	pb pb.WalletServiceClient
}

func NewGrpcWalletClient(pb pb.WalletServiceClient) port.WalletService {
	return &walletRpcCli{pb: pb}
}

func (s *walletRpcCli) GetUserWallet(ctx context.Context, userUUID uuid.UUID) (*entity.UserWallet, error) {
	req := pb.UserWalletRequest{
		UserUuid: userUUID.String(),
	}
	resp, err := s.pb.GetUserWallet(ctx, &req)
	if err != nil {
		return nil, err
	}

	respUserUUID, err := uuid.Parse(resp.UserUuid)
	if err != nil {
		return nil, err
	}

	wallet := &entity.UserWallet{
		UserUUID: respUserUUID,
		Fund:     resp.Fund,
	}

	return wallet, nil

}

func (s *walletRpcCli) IncreaseUserWallet(ctx context.Context, userUUID uuid.UUID, fund int64) (*entity.UserWallet, error) {
	req := pb.UpdateWalletRequest{
		UserUuid: userUUID.String(),
		Fund:     fund,
	}
	resp, err := s.pb.IncreaseUserWallet(ctx, &req)
	if err != nil {
		return nil, err
	}

	respUserUUID, err := uuid.Parse(resp.UserUuid)
	if err != nil {
		return nil, err
	}

	wallet := &entity.UserWallet{
		UserUUID: respUserUUID,
		Fund:     resp.Fund,
	}

	return wallet, nil
}

func (s *walletRpcCli) DecreaseUserWallet(ctx context.Context, userUUID uuid.UUID, fund int64) (*entity.UserWallet, error) {
	req := pb.UpdateWalletRequest{
		UserUuid: userUUID.String(),
		Fund:     fund,
	}
	resp, err := s.pb.DecreaseUserWallet(ctx, &req)
	if err != nil {
		return nil, err
	}

	respUserUUID, err := uuid.Parse(resp.UserUuid)
	if err != nil {
		return nil, err
	}

	wallet := &entity.UserWallet{
		UserUUID: respUserUUID,
		Fund:     resp.Fund,
	}

	return wallet, nil
}
