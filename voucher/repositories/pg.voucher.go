package repository

import (
	"github.com/jmoiron/sqlx"
	entity "github.com/kaolnwza/muniverse/voucher/entities"
)

func GetVoucherByCode(q sqlx.Queryer, dest *entity.Voucher, code string) error {
	query := `
		SELECT 
			uuid,
			voucher_name,
			voucher_code,
			discount_type,
			discount,
			max_discount,
			voucher_quantity,
			voucher_status,
			expired_at
		FROM voucher
		WHERE voucher_code = $1
		AND expired_at > now()
		AND voucher_status IS TRUE
	`

	return sqlx.Get(q, dest, query, code)
}
