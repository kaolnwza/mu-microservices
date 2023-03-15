package rpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/kaolnwza/muniverse/seer/internal/adapters/grpc/proto/pb"
	port "github.com/kaolnwza/muniverse/seer/internal/ports"
)

type upiGrpc struct {
	pb pb.ProfileServiceClient
}

func NewGrpcStorerClient(pb pb.ProfileServiceClient) port.StorageService {
	return upiGrpc{pb: pb}
}

func (r upiGrpc) GetUserProfileImage(userUUID uuid.UUID) (*string, error) {
	resp, err := r.pb.GetProfileImage(context.Background(), &pb.ProfileRequest{UserUuid: userUUID.String()})
	if err != nil {
		return nil, err
	}

	// r.pb.GetProfileImage()

	return &resp.Url, nil
}
