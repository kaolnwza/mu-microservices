package rpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/kaolnwza/muniverse/order/internal/adapters/grpc/proto/pb"
	entity "github.com/kaolnwza/muniverse/order/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/order/internal/ports"
)

type seerRpcCli struct {
	pb pb.SeerServiceClient
}

func NewGrpcSeerClient(pb pb.SeerServiceClient) port.SeerService {
	return &seerRpcCli{pb: pb}
}

func (r *seerRpcCli) GetSeerByUserUUID(ctx context.Context, userUUID uuid.UUID) (*entity.SeerResponse, error) {
	resp, err := r.pb.GetSeerByUserUUID(ctx, &pb.SeerRequest{UserUuid: userUUID.String()})
	if err != nil {
		return nil, err
	}

	seerUUID, _ := uuid.Parse(resp.Uuid)
	user := &entity.SeerResponse{
		UUID:               seerUUID,
		OnsiteAvailable:    resp.OnsiteAvailable,
		ChatAvailable:      resp.ChatAvailable,
		CallAvailable:      resp.CallAvailable,
		VideoCallAvailable: resp.VideoCallAvailable,
		Major:              resp.Major,
		MajorDescription:   resp.MajorDescription,
		DescriptionProfile: resp.DescriptionProfile,
		MapCoordinate:      resp.MapCoordinate,
		ImageURL:           resp.ImageUrl,
		DisplayName:        resp.DisplayName,
	}

	return user, nil
}

func (r *seerRpcCli) GetUserUUIDBySeerUUID(ctx context.Context, seerUUID uuid.UUID) (*uuid.UUID, error) {
	resp, err := r.pb.GetUserUUIDBySeerUUID(ctx, &pb.SeerUUIDRequest{SeerUuid: seerUUID.String()})
	if err != nil {
		return nil, err
	}

	userUUID, err := uuid.Parse(resp.UserUuid)
	if err != nil {
		return nil, err
	}

	return &userUUID, nil
}
