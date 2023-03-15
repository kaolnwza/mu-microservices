package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kaolnwza/muniverse/storage/entity"
)

func FetchImageProfile(tx sqlx.Queryer, upload *entity.Upload, userUUID uuid.UUID) error {
	s := `
		SELECT bucket, path
		FROM upload
		WHERE 
			EXISTS (
				SELECT 1
				FROM profile_image
				WHERE upload_uuid = upload.uuid
				AND user_uuid = $1
			)
			AND deleted_at IS NULL
	`

	return sqlx.Get(tx, upload, s, userUUID)
}

func NewImageProfile(tx sqlx.Execer, userUUID uuid.UUID, uploadUUID uuid.UUID) error {
	query := `
		INSERT INTO profile_image (user_uuid, upload_uuid)
		VALUES ($1, $2)
		ON CONFLICT (user_uuid)
		DO UPDATE
		SET
			upload_uuid = $2
	`

	_, err := tx.Exec(query, userUUID, uploadUUID)
	return err
}
