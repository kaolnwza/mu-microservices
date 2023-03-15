package port

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/order/internal/application/core/entities"
)

type HoroOrderRepository interface {
	Create(ctx context.Context, userUUID uuid.UUID, horoUUID uuid.UUID, paymentUUID uuid.UUID, voucherUUID string, price int, startTime string, endTime string) error
	GetUpcomingCustomerOrder(ctx context.Context, dest *[]*entity.OrderCustomer, seerUUID uuid.UUID) error
	GetCustomerOrderHistory(ctx context.Context, dest *[]*entity.OrderCustomer, seerUUID uuid.UUID) error
	GetOrderByUserUUID(ctx context.Context, dest *[]*entity.OrderCustomer, userUUID uuid.UUID) error
	GetOrderHistoryByUserUUID(ctx context.Context, dest *[]*entity.OrderCustomer, userUUID uuid.UUID) error
}

type HoroOrderService interface {
	CreateHoroOrder(ctx context.Context, userUUID uuid.UUID, horoUUID uuid.UUID, voucherUUID uuid.UUID, price int, startTime string, endTime string) error
	GetUpcomingCustomerOrder(ctx context.Context, seerUUID uuid.UUID) (*[]*entity.OrderCustomer, error)
	GetCustomerOrderHistory(ctx context.Context, seerUUID uuid.UUID) (*[]*entity.OrderCustomer, error)
	GetOrderByUserUUID(ctx context.Context, userUUID uuid.UUID) (*[]*entity.OrderCustomer, error)
	GetOrderHistoryByUserUUID(ctx context.Context, userUUID uuid.UUID) (*[]*entity.OrderCustomer, error)
}
