package rpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/kaolnwza/muniverse/wallet/internal/adapters/grpc/proto/pb"
	port "github.com/kaolnwza/muniverse/wallet/internal/ports"
)

type rpcWalletSrv struct {
	pb.WalletServiceServer
	svc port.WalletService
}

func NewRpcWalletServer(svc port.WalletService) pb.WalletServiceServer {
	return &rpcWalletSrv{svc: svc}
}

func (r *rpcWalletSrv) CreateUserWallet(ctx context.Context, req *pb.UserWalletRequest) (*pb.EmptyResponse, error) {
	userUUID, err := uuid.Parse(req.UserUuid)
	if err != nil {
		return nil, err
	}

	if err := r.svc.CreateUserWallet(ctx, userUUID); err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *rpcWalletSrv) GetUserWallet(ctx context.Context, req *pb.UserWalletRequest) (*pb.WalletResponse, error) {
	userUUID, err := uuid.Parse(req.UserUuid)
	if err != nil {
		return nil, err
	}

	wallet, err := r.svc.GetWalletByUserUUID(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	return &pb.WalletResponse{
		UserUuid: req.UserUuid,
		Fund:     wallet.Fund,
	}, nil
}

func (r *rpcWalletSrv) IncreaseUserWallet(ctx context.Context, req *pb.UpdateWalletRequest) (*pb.WalletResponse, error) {
	userUUID, err := uuid.Parse(req.UserUuid)
	if err != nil {
		return nil, err
	}

	wallet, err := r.svc.IncreaseWallet(ctx, userUUID, req.Fund)
	if err != nil {
		return nil, err
	}

	return &pb.WalletResponse{
		UserUuid: req.UserUuid,
		Fund:     wallet.Fund,
	}, nil
}

func (r *rpcWalletSrv) DecreaseUserWallet(ctx context.Context, req *pb.UpdateWalletRequest) (*pb.WalletResponse, error) {
	userUUID, err := uuid.Parse(req.UserUuid)
	if err != nil {
		return nil, err
	}

	wallet, err := r.svc.DecreaseWallet(ctx, userUUID, req.Fund)
	if err != nil {
		return nil, err
	}

	return &pb.WalletResponse{
		UserUuid: req.UserUuid,
		Fund:     wallet.Fund,
	}, nil
}
