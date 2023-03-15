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

type storageRpcServer struct {
	pb.StorageServiceServer
}

func NewStorageRpcServer() pb.StorageServiceServer {
	return &storageRpcServer{}
}

func (r *storageRpcServer) GetImageByUploadUUID(ctx context.Context, req *pb.GetImageURLRequest) (*pb.ImageURLResponse, error) {
	upload := entity.Upload{}
	uploadUUID, _ := uuid.Parse(req.UploadUuid)

	err := repository.GetByUploadUUID(database.NewPostgresDB(), &upload, uploadUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			return &pb.ImageURLResponse{Url: ""}, nil
		}

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

	return &pb.ImageURLResponse{Url: url}, nil
}
