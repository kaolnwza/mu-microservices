package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kaolnwza/muniverse/storage/entity"
)

func UploadImage(tx sqlx.Queryer, upload *entity.Upload, path string, bucket string, userUUID uuid.UUID) error {
	s := `
		INSERT INTO upload (user_uuid, bucket, path)
		SELECT $1, $2, $3
		RETURNING uuid
	`

	return sqlx.Get(tx, upload, s, userUUID, bucket, path)

}

func GetByUploadUUID(tx sqlx.Queryer, upload *entity.Upload, uploadUUID uuid.UUID) error {
	query := `
		SELECT bucket, path
		FROM upload
		WHERE uuid = $1
	`

	return sqlx.Get(tx, upload, query, uploadUUID)
}
