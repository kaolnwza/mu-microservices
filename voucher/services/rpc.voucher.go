package service

import (
	"context"

	"github.com/kaolnwza/muniverse/voucher/database"
	entity "github.com/kaolnwza/muniverse/voucher/entities"
	pb "github.com/kaolnwza/muniverse/voucher/proto/pb"
	repository "github.com/kaolnwza/muniverse/voucher/repositories"
)

type RpcVoucherServer struct {
	pb.VoucherServiceServer
}

func (r *RpcVoucherServer) ValidateVoucherCode(ctx context.Context, req *pb.VoucherCodeRequest) (*pb.VoucherStatusResponse, error) {
	return nil, nil
}

func (r *RpcVoucherServer) GetVoucherByCode(ctx context.Context, req *pb.VoucherCodeRequest) (*pb.VoucherResponse, error) {
	conn := database.NewPostgresDB()
	voucher := entity.Voucher{}
	if err := repository.GetVoucherByCode(conn, &voucher, req.VoucherCode); err != nil {
		return nil, err
	}

	resp := pb.VoucherResponse{
		VoucherName:     voucher.VoucherName,
		DiscountType:    voucher.DiscountType,
		Discount:        int32(voucher.Discount),
		VoucherQuantity: int32(voucher.VoucherQuantity),
		ExpiredAt:       voucher.ExpiredAt.String(),
		Status:          voucher.VoucherStatus,
	}

	return &resp, nil
}
