package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kaolnwza/muniverse/storage/entity"
)

func NewHoroServiceImage(tx sqlx.Execer, horoUUID uuid.UUID, val string) error {
	query := `
		WITH del_old AS (
			DELETE FROM horo_service_img
			WHERE horo_service_uuid = $1
		)

		INSERT INTO horo_service_img (horo_service_uuid, upload_uuid, image_order)
		VALUES ` + val

	_, err := tx.Exec(query, horoUUID)
	return err
}

func FetchHoroServiceImage(tx sqlx.Queryer, dest *[]*entity.HoroServiceImage, horoUUID uuid.UUID) error {
	query := `
	SELECT bucket, path, image_order
	FROM upload
	LEFT JOIN horo_service_img ON upload.uuid = upload_uuid
	WHERE
		horo_service_uuid = $1
		AND deleted_at IS NULL
	ORDER BY image_order
	`

	return sqlx.Select(tx, dest, query, horoUUID)

}
