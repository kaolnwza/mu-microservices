package entity

import (
	"time"

	"github.com/google/uuid"
)

// 'fail', 'unpaid', 'paid', 'cancel_by_user', 'cancel_by_seer', 'success', 'inprogress', 'refund'

type HoroOrderStatus string

const (
	HORO_ORDER_STATUS_FAIL    HoroOrderStatus = "fail"
	HORO_ORDER_STATUS_UNPAID  HoroOrderStatus = "unpaid"
	HORO_ORDER_STATUS_PAID    HoroOrderStatus = "paid"
	HORO_ORDER_STATUS_SUCCESS HoroOrderStatus = "success"
)

type Order struct {
	UUID            uuid.UUID       `db:"uuid"`
	UserUUID        uuid.UUID       `db:"user_uuid"`
	HoroServiceUUID uuid.UUID       `db:"horo_service_uuid"`
	VoucerUUID      uuid.NullUUID   `db:"voucher_uuid"`
	PaymentUUID     uuid.UUID       `db:"payment_uuid"`
	Price           int             `db:"price"`
	Status          HoroOrderStatus `db:"status"`
	StartTime       time.Time       `db:"start_time"`
	EndTime         time.Time       `db:"end_time"`
}

type OrderCustomer struct {
	UUID            uuid.UUID `db:"uuid" json:"uuid"`
	UserUUID        uuid.UUID `db:"user_uuid" json:"user_uuid"`
	HoroServiceUUID uuid.UUID `db:"horo_service_uuid" json:"horo_service_uuid"`
	StartTime       time.Time `db:"start_time" json:"start_time"`
	EndTime         time.Time `db:"end_time" json:"end_time"`
}
