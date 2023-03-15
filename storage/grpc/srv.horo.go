package rpc

import (
	"context"
	"fmt"
	"io"

	"github.com/google/uuid"
	"github.com/kaolnwza/muniverse/storage/database"
	"github.com/kaolnwza/muniverse/storage/entity"
	"github.com/kaolnwza/muniverse/storage/helper"
	"github.com/kaolnwza/muniverse/storage/proto/pb"
	repository "github.com/kaolnwza/muniverse/storage/repositories"
	"google.golang.org/protobuf/types/known/emptypb"
)

type horoRpcServer struct {
	pb.HoroSvcServiceServer
}

func NewHoroRpcService() pb.HoroSvcServiceServer {
	return &horoRpcServer{}
}

func (r *horoRpcServer) NewHoroServiceImage(stream pb.HoroSvcService_NewHoroServiceImageServer) error {
	val := []string{}
	var horoUUID uuid.UUID
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			queryValues := ``
			for idx, i := range val {
				queryValues += i
				if idx != len(val)-1 {
					queryValues += ","
				}
			}

			if err := repository.NewHoroServiceImage(database.NewPostgresDB(), horoUUID, queryValues); err != nil {
				return err
			}

			return stream.SendAndClose(&emptypb.Empty{})
		}

		if err != nil {
			return err
		}

		horoUUID, err = uuid.Parse(req.HoroServiceUuid)
		if err != nil {
			return err
		}

		val = append(val, fmt.Sprintf(`('%s', '%s', %v)`, req.HoroServiceUuid, req.UploadUuid, req.Order))
	}

}

func (r *horoRpcServer) GetHoroServiceImage(req *pb.GetHoroServiceRequest, stream pb.HoroSvcService_GetHoroServiceImageServer) error {
	horoUUID, err := uuid.Parse(req.HoroServiceUuid)
	if err != nil {
		return err
	}

	images := []*entity.HoroServiceImage{}
	if err := repository.FetchHoroServiceImage(database.NewPostgresDB(), &images, horoUUID); err != nil {
		return err
	}

	gcsCh := make(chan error)
	streamErr := make(chan error)
	for _, item := range images {
		go func(item *entity.HoroServiceImage) {
			url, err := helper.GenerateImageURI(context.Background(), item.Bucket, item.Path)
			gcsCh <- err

			streamErr <- stream.Send(&pb.HoroServiceImageResponse{
				Order: item.ImageOrder,
				Url:   url,
			})
		}(item)
	}

	for i := 0; i < len(images); i++ {
		if err := <-gcsCh; err != nil {
			return err
		}

		if err := <-streamErr; err != nil {
			return err
		}
	}

	return nil
}
