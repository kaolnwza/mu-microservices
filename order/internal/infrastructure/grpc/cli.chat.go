package rpc

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kaolnwza/muniverse/order/internal/adapters/grpc/proto/pb"
	port "github.com/kaolnwza/muniverse/order/internal/ports"
)

type chatRpcCli struct {
	pb pb.ChatServiceClient
}

func NewChatServiceClient(pb pb.ChatServiceClient) port.ChatService {
	return &chatRpcCli{pb: pb}
}

func (r *chatRpcCli) CreateChatRoom(ctx context.Context, orderUUID uuid.UUID, userUUID uuid.UUID, seerUUID uuid.UUID, startTime time.Time, endTime time.Time) error {
	req := pb.CreateChatRequest{
		StUserUuid: seerUUID.String(),
		NdUserUuid: userUUID.String(),
		StartTime:  startTime.Format(time.RFC3339),
		EndTime:    endTime.Format(time.RFC3339),
		OrderUuid:  orderUUID.String(),
	}

	_, err := r.pb.CreateChat(ctx, &req)
	if err != nil {
		return err
	}

	return nil

}
