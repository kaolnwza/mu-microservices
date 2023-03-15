package pgsql

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/seer/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/seer/internal/ports"
)

type seerRepo struct {
	tx port.Transactor
}

func NewSeerRepository(tx port.Transactor) port.SeerRepository {
	return &seerRepo{tx: tx}
}

func (r *seerRepo) GetByUUID(ctx context.Context, seer *entity.Seer, seerUUID uuid.UUID) error {
	query := `
		SELECT
			uuid,
			user_uuid,
			onsite_available,
			chat_available,
			call_available,
			video_call_available,
			major,
			major_description,
			description_profile,
			map_coordinate
		FROM seer
		WHERE uuid = $1
	`

	return r.tx.GetOne(ctx, seer, query, seerUUID)
}

func (r *seerRepo) GetByUserUUID(ctx context.Context, seer *entity.Seer, userUUID uuid.UUID) error {
	query := `
		SELECT
			uuid,
			user_uuid,
			onsite_available,
			chat_available,
			call_available,
			video_call_available,
			major,
			major_description,
			description_profile,
			map_coordinate
		FROM seer
		WHERE user_uuid = $1
	`

	return r.tx.GetOne(ctx, seer, query, userUUID)
}
