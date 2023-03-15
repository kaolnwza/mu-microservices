package rpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/kaolnwza/muniverse/feed/internal/adapters/grpc/proto/pb"
	port "github.com/kaolnwza/muniverse/feed/internal/ports"
)

type storageRpcCli struct {
	pb pb.StorageServiceClient
}

func NewGrpcStorerClient(pb pb.StorageServiceClient) port.ImageStorer {
	return storageRpcCli{pb: pb}
}

func (r storageRpcCli) GetURLByUploadUUID(ctx context.Context, uploadUUID uuid.UUID) (*string, error) {
	if uploadUUID == uuid.Nil {
		return nil, nil
	}

	resp, err := r.pb.GetImageByUploadUUID(ctx, &pb.GetImageURLRequest{UploadUuid: uploadUUID.String()})
	if err != nil {
		return nil, err
	}

	return &resp.Url, nil
}
