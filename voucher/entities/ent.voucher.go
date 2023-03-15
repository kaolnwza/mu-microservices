package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Voucher struct {
	UUID            uuid.UUID     `db:"uuid"`
	VoucherName     string        `db:"voucher_name"`
	VoucherCode     string        `db:"voucher_code"`
	DiscountType    string        `db:"discount_type"`
	Discount        int           `db:"discount"`
	MaxDiscount     sql.NullInt16 `db:"max_discount"`
	VoucherQuantity int           `db:"voucher_quantity"`
	VoucherStatus   bool          `db:"voucher_status"`
	CreatedAt       time.Time     `db:"-"`
	ExpiredAt       time.Time     `db:"expired_at"`
}
