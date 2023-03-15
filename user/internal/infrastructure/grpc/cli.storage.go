package rpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/kaolnwza/muniverse/user/internal/adapters/grpc/proto/pb"
	port "github.com/kaolnwza/muniverse/user/internal/ports"
)

type upiGrpc struct {
	pb pb.ProfileServiceClient
}

func NewGrpcStorerClient(pb pb.ProfileServiceClient) port.ImageStorer {
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

func (r upiGrpc) UpdateUserProfileImage(userUUID uuid.UUID, uploadUUID uuid.UUID) error {
	_, err := r.pb.NewProfileImage(context.Background(), &pb.NewProfileRequest{
		UserUuid:   userUUID.String(),
		UploadUuid: uploadUUID.String(),
	})

	if err != nil {
		return err
	}

	return nil
}
