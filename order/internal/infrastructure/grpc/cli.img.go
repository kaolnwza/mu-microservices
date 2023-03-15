package rpc

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/kaolnwza/muniverse/order/internal/adapters/grpc/proto/pb"
	entity "github.com/kaolnwza/muniverse/order/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/order/internal/ports"
)

type imgRpcCli struct {
	pb pb.HoroSvcServiceClient
}

func NewImageStorageServiceClient(pb pb.HoroSvcServiceClient) port.ImageStorageService {
	return &imgRpcCli{pb: pb}
}

func (r *imgRpcCli) GetHoroServiceImages(ctx context.Context, horoUUID uuid.UUID) (*[]*entity.HoroImages, error) {
	req := pb.GetHoroServiceRequest{HoroServiceUuid: horoUUID.String()}
	stream, err := r.pb.GetHoroServiceImage(ctx, &req)
	if err != nil {
		return nil, err
	}

	temp := make([]*entity.HoroImages, 10)
	img := []*entity.HoroImages{}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		temp[int(resp.Order)-1] = &entity.HoroImages{
			Url:        resp.Url,
			ImageOrder: resp.Order,
		}

	}

	for _, item := range temp {
		if item != nil {
			img = append(img, item)
		}
	}

	return &img, nil
}

func (r *imgRpcCli) NewHoroServiceImages(ctx context.Context, horoUUID uuid.UUID, images []*entity.HoroImagesRequest) error {
	stream, err := r.pb.NewHoroServiceImage(ctx)
	if err != nil {
		return err
	}

	for _, item := range images {
		sender := &pb.NewHoroServiceRequest{
			HoroServiceUuid: horoUUID.String(),
			UploadUuid:      item.UploadUUID.String(),
			Order:           item.ImageOrder,
		}

		if err := stream.Send(sender); err != nil {
			if err := stream.CloseSend(); err != nil {
				return err
			}

			return err
		}
	}

	if _, err := stream.CloseAndRecv(); err != nil {
		return err
	}

	return nil
}
