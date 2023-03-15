package rpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/kaolnwza/muniverse/user/internal/adapters/grpc/proto/pb"
	port "github.com/kaolnwza/muniverse/user/internal/ports"
)

type walRpcCli struct {
	pb pb.WalletServiceClient
}

func NewGrpcWalletClient(pb pb.WalletServiceClient) port.WalletService {
	return &walRpcCli{pb: pb}
}

func (s *walRpcCli) CreateNewUserWallet(ctx context.Context, userUUID uuid.UUID) error {
	req := &pb.UserWalletRequest{
		UserUuid: userUUID.String(),
	}

	if _, err := s.pb.CreateUserWallet(ctx, req); err != nil {
		return err
	}

	return nil
}
