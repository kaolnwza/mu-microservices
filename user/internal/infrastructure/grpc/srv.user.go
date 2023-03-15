package rpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/kaolnwza/muniverse/user/internal/adapters/grpc/proto/pb"
	port "github.com/kaolnwza/muniverse/user/internal/ports"
	log "github.com/kaolnwza/muniverse/user/lib/logs"
)

type rpcUserServer struct {
	pb.UserServiceServer
	svc port.UserService
}

func NewGrpcUserServer(svc port.UserService) pb.UserServiceServer {
	return &rpcUserServer{svc: svc}
}

func (s *rpcUserServer) GetUserByUUID(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	userUUID, _ := uuid.Parse(req.Uuid)
	user, err := s.svc.GetUserWithoutImgByUUID(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	resp := pb.UserResponse{
		Uuid:              user.UUID.String(),
		DisplayName:       user.DisplayName,
		Birthday:          user.Birthday,
		Description:       user.Description,
		TelNumber:         user.TelNumber,
		Role:              string(user.Role),
		ProfilePictureUrl: user.ProfilePictureURL,
	}
	log.Info("passall")
	return &resp, nil
}

func (s *rpcUserServer) GetUserWithoutImgByUUID(ctx context.Context, req *pb.UserRequest) (*pb.UserWithoutImgResponse, error) {
	userUUID, _ := uuid.Parse(req.Uuid)
	user, err := s.svc.GetUserByUUID(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	resp := pb.UserWithoutImgResponse{
		Uuid:        user.UUID.String(),
		DisplayName: user.DisplayName,
		Birthday:    user.Birthday,
		Description: user.Description,
		TelNumber:   user.TelNumber,
		Role:        string(user.Role),
	}

	return &resp, nil
}

func (s *rpcUserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := s.svc.CreateUser(ctx, req.DisplayName, req.Birthday, req.Description)
	if err != nil {
		return nil, err
	}

	resp := &pb.CreateUserResponse{
		Uuid:        user.UUID.String(),
		DisplayName: user.DisplayName,
		Birthday:    user.Birthday.Time.String(),
		Description: user.Description.String,
	}

	return resp, nil
}
