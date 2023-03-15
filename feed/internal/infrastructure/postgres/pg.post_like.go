package postgres

import (
	"context"

	"github.com/google/uuid"
	port "github.com/kaolnwza/muniverse/feed/internal/ports"
)

type likeRepo struct {
	tx port.Transactor
}

func NewLikeRepository(tx port.Transactor) port.LikeRepository {
	return &likeRepo{tx: tx}
}

func (r *likeRepo) PostLike(ctx context.Context, status *bool, userUUID uuid.UUID, postUUID uuid.UUID) error {
	query := `
		INSERT INTO post_like (user_uuid, post_uuid)
		VALUES ($1, $2)
		ON CONFLICT
		DO NOTHING
	`

	return r.tx.Insert(ctx, query, userUUID, postUUID)
}

func (r *likeRepo) PostUnlike(ctx context.Context, status *bool, userUUID uuid.UUID, postUUID uuid.UUID) error {
	query := `
		DELETE FROM post_like
		WHERE user_uuid = $1
		AND post_uuid = $2
	`

	return r.tx.Delete(ctx, query, userUUID, postUUID)
}
