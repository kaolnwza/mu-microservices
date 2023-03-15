package rpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/kaolnwza/muniverse/seer/internal/adapters/grpc/proto/pb"
	entity "github.com/kaolnwza/muniverse/seer/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/seer/internal/ports"
)

type userRpcCli struct {
	pb pb.UserServiceClient
}

func NewGrpcUserClient(pb pb.UserServiceClient) port.UserService {
	return userRpcCli{pb: pb}
}

func (r userRpcCli) GetUserInfo(ctx context.Context, userUUID uuid.UUID) (*entity.UserResponse, error) {
	resp, err := r.pb.GetUserWithoutImgByUUID(ctx, &pb.UserRequest{Uuid: userUUID.String()})
	if err != nil {
		return nil, err
	}

	user := &entity.UserResponse{
		UUID:        userUUID,
		DisplayName: resp.DisplayName,
		Birthday:    resp.Birthday,
		Description: resp.Description,
		TelNumber:   resp.TelNumber,
		Email:       resp.Email,
		Role:        resp.Role,
		// ProfilePictureURL: resp.ProfilePictureUrl,
	}
	// r.pb.GetProfileImage()

	return user, nil
}
