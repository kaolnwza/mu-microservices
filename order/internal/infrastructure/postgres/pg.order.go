package postgres

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/order/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/order/internal/ports"
)

type horoOdrRepo struct {
	tx port.Transactor
}

func NewHoroOrderRepository(tx port.Transactor) port.HoroOrderRepository {
	return &horoOdrRepo{tx: tx}
}

func (h *horoOdrRepo) Create(ctx context.Context, userUUID uuid.UUID, horoUUID uuid.UUID, paymentUUID uuid.UUID, voucherUUID string, price int, startTime string, endTime string) error {
	query := `
		INSERT INTO horo_order (user_uuid, horo_service_uuid, voucher_uuid, payment_uuid, price, start_time, end_time)
		VALUES ($1, $2, ` + voucherUUID + `, $3, $4, $5, $6)
	`

	return h.tx.Insert(ctx, query, userUUID, horoUUID, paymentUUID, price, startTime, endTime)
}

func (h *horoOdrRepo) GetUpcomingCustomerOrder(ctx context.Context, dest *[]*entity.OrderCustomer, seerUUID uuid.UUID) error {
	query := `
		SELECT
			uuid,
			user_uuid,
			horo_service_uuid,
			status,
			start_time,
			end_time
		FROM horo_order
		WHERE EXISTS (
			SELECT 1 
			FROM horo_service
			WHERE horo_service.uuid = horo_order.horo_service_uuid
			AND seer_uuid = $1
		)
		AND status IN ('paid', 'confirmed')
	`

	return h.tx.Get(ctx, dest, query, seerUUID)
}

func (h *horoOdrRepo) GetCustomerOrderHistory(ctx context.Context, dest *[]*entity.OrderCustomer, seerUUID uuid.UUID) error {
	query := `
		SELECT
			uuid,
			user_uuid,
			horo_service_uuid,
			status,
			start_time,
			end_time
		FROM horo_order
		WHERE EXISTS (
			SELECT 1 
			FROM horo_service
			WHERE horo_service.uuid = horo_order.horo_service_uuid
			AND seer_uuid = $1
		)
		AND status IN ('success')
	`

	return h.tx.Get(ctx, dest, query, seerUUID)
}

func (h *horoOdrRepo) GetOrderByUserUUID(ctx context.Context, dest *[]*entity.OrderCustomer, userUUID uuid.UUID) error {
	query := `
		SELECT
			uuid,
			user_uuid,
			horo_service_uuid,
			status,
			start_time,
			end_time
		FROM horo_order
		WHERE status IN ('paid', 'confirmed')
		AND user_uuid = $1
	`

	return h.tx.Get(ctx, dest, query, userUUID)
}

func (h *horoOdrRepo) GetOrderHistoryByUserUUID(ctx context.Context, dest *[]*entity.OrderCustomer, userUUID uuid.UUID) error {
	query := `
		SELECT
			uuid,
			user_uuid,
			horo_service_uuid,
			status,
			start_time,
			end_time
		FROM horo_order
		WHERE status IN ('success')
		AND user_uuid = $1
	`

	return h.tx.Get(ctx, dest, query, userUUID)
}

func (h *horoOdrRepo) UpdateOrderStatusByUUID(ctx context.Context, status entity.HoroOrderStatus, horoUUID uuid.UUID) error {
	query := `
		UPDATE horo_order
		SET status = $1
		WHERE uuid = $2
	`

	return h.tx.Update(ctx, query, status, horoUUID)
}

func (h *horoOdrRepo) GetOrderByUUID(ctx context.Context, dest *entity.Order, orderUUID uuid.UUID) error {
	query := `
		SELECT
			uuid,
			user_uuid,
			horo_service_uuid,
			status,
			start_time,
			end_time
		FROM horo_order
		WHERE uuid = $1
	`

	return h.tx.GetOne(ctx, dest, query, orderUUID)
}
