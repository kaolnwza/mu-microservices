package rpc

import (
	"context"

	"github.com/kaolnwza/muniverse/order/internal/adapters/grpc/proto/pb"
	entity "github.com/kaolnwza/muniverse/order/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/order/internal/ports"
)

type voucherRpcCli struct {
	pb pb.VoucherServiceClient
}

func NewGrpcVoucherClient(pb pb.VoucherServiceClient) port.VoucherService {
	return &voucherRpcCli{pb: pb}
}

func (r *voucherRpcCli) GetVoucherByCode(ctx context.Context, code string) (*entity.Voucher, error) {
	req := pb.VoucherCodeRequest{VoucherCode: code}
	resp, err := r.pb.GetVoucherByCode(ctx, &req)
	if err != nil {
		return nil, err
	}

	voucher := entity.Voucher{
		VoucherName:     resp.VoucherName,
		DiscountType:    resp.DiscountType,
		Discount:        int(resp.Discount),
		VoucherQuantity: int(resp.VoucherQuantity),
		VoucherStatus:   resp.Status,
	}

	return &voucher, nil

}
