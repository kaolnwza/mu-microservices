package rpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/kaolnwza/muniverse/chat/internal/adapters/grpc/proto/pb"
	port "github.com/kaolnwza/muniverse/chat/internal/ports"
	"google.golang.org/protobuf/types/known/emptypb"
)

type rpcChatServer struct {
	pb.ChatServiceServer
	roomSvc port.RoomService
}

func NewGrpcChatServer(roomSvc port.RoomService) pb.ChatServiceServer {
	return &rpcChatServer{roomSvc: roomSvc}
}

func (r *rpcChatServer) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*emptypb.Empty, error) {
	stUserUUID, err := uuid.Parse(req.StUserUuid)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	ndUserUUID, err := uuid.Parse(req.NdUserUuid)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	orderUUID, err := uuid.Parse(req.OrderUuid)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	if err := r.roomSvc.CreateRoom(ctx, stUserUUID, ndUserUUID, req.StartTime, req.EndTime, orderUUID); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}

func (r *rpcChatServer) UpdateChatRoomStatus(ctx context.Context, req *pb.ChatRoomStatusRequest) (*emptypb.Empty, error) {
	orderUUID, err := uuid.Parse(req.HoroOrderUuid)
	if err != nil {
		return nil, err
	}

	if err := r.roomSvc.UpdateRoomStatusByOrderUUID(ctx, orderUUID, req.Status); err != nil {
		return nil, err
	}

	return nil, nil
}
