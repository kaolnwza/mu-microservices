package rpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/kaolnwza/muniverse/seer/internal/adapters/grpc/proto/pb"
	port "github.com/kaolnwza/muniverse/seer/internal/ports"
)

type rpcSeerServer struct {
	pb.SeerServiceServer
	svc port.SeerService
}

func NewGrpcSeerServer(svc port.SeerService) pb.SeerServiceServer {
	return &rpcSeerServer{svc: svc}
}

func (r *rpcSeerServer) GetSeerByUserUUID(ctx context.Context, req *pb.SeerRequest) (*pb.SeerResponse, error) {
	userUUID, err := uuid.Parse(req.UserUuid)
	if err != nil {
		return nil, err
	}

	seer, err := r.svc.GetSeerByUserUUID(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	resp := pb.SeerResponse{
		Uuid:               seer.UUID.String(),
		OnsiteAvailable:    seer.OnsiteAvailable,
		ChatAvailable:      seer.ChatAvailable,
		CallAvailable:      seer.CallAvailable,
		VideoCallAvailable: seer.VideoCallAvailable,
		Major:              seer.Major,
		MajorDescription:   seer.MajorDescription,
		DescriptionProfile: seer.DescriptionProfile,
		MapCoordinate:      seer.MapCoordinate,
		ImageUrl:           seer.ImageURL,
		DisplayName:        seer.DisplayName,
	}

	return &resp, err
}

func (r *rpcSeerServer) GetUserUUIDBySeerUUID(ctx context.Context, req *pb.SeerUUIDRequest) (*pb.UserUUIDResponse, error) {
	seerUUID, err := uuid.Parse(req.SeerUuid)
	if err != nil {
		return nil, err
	}

	seer, err := r.svc.GetSeerByUUID(ctx, seerUUID)
	if err != nil {
		return nil, err
	}

	return &pb.UserUUIDResponse{
		UserUuid: seer.UserUUID.String(),
	}, nil
}
