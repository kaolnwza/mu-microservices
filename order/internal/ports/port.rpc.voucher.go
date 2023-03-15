package port

import (
	"context"

	entity "github.com/kaolnwza/muniverse/order/internal/application/core/entities"
)

type VoucherService interface {
	GetVoucherByCode(ctx context.Context, code string) (*entity.Voucher, error)
}
