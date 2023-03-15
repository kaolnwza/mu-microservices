package rpc

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/kaolnwza/muniverse/storage/database"
	"github.com/kaolnwza/muniverse/storage/entity"
	"github.com/kaolnwza/muniverse/storage/helper"
	log "github.com/kaolnwza/muniverse/storage/logs"
	"github.com/kaolnwza/muniverse/storage/proto/pb"
	repository "github.com/kaolnwza/muniverse/storage/repositories"
)

type profileRpcServer struct {
	pb.ProfileServiceServer
}

func NewProfileRpcServer() pb.ProfileServiceServer {
	return &profileRpcServer{}
}

func (s *profileRpcServer) GetProfileImage(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileResponse, error) {
	upload := entity.Upload{}
	userUUID, _ := uuid.Parse(req.UserUuid)

	err := repository.FetchImageProfile(database.NewPostgresDB(), &upload, userUUID)
	if err != nil && err != sql.ErrNoRows {
		log.Error(err)
		return nil, err
	}

	url := ""
	if err != sql.ErrNoRows {
		url, err = helper.GenerateImageURI(ctx, upload.Bucket, upload.Path)
		if err != nil {
			log.Error(err)
			return nil, err
		}
	}

	return &pb.ProfileResponse{Url: url}, nil
}

func (s *profileRpcServer) NewProfileImage(ctx context.Context, req *pb.NewProfileRequest) (*pb.NewProfileResponse, error) {
	uploadUUID, _ := uuid.Parse(req.UploadUuid)
	userUUID, _ := uuid.Parse(req.UserUuid)

	if err := repository.NewImageProfile(database.NewPostgresDB(), userUUID, uploadUUID); err != nil {
		return nil, err
	}

	return &pb.NewProfileResponse{}, nil
}
