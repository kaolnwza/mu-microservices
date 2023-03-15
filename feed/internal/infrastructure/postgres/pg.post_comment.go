package postgres

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/feed/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/feed/internal/ports"
)

type comntRepo struct {
	tx port.Transactor
}

func NewCommentRepository(tx port.Transactor) port.CommentRepository {
	return &comntRepo{tx: tx}
}

func (r *comntRepo) CreateComment(ctx context.Context, comment *entity.Comment, userUUID uuid.UUID, postUUID uuid.UUID, text string) error {
	query := `
		INSERT INTO post_comment (user_uuid, post_uuid, comment)
		VALUES ($1, $2, $3)
		RETURNING uuid, user_uuid, post_uuid, comment, created_at
	`

	return r.tx.InsertWithReturningOne(ctx, comment, query, userUUID, postUUID, text)
}

func (r *comntRepo) GetCommentByPostUUID(ctx context.Context, comment *[]*entity.Comment, postUUID uuid.UUID) error {
	query := `
		SELECT
			uuid,
			user_uuid,
			post_uuid,
			comment,
			created_at,
			deleted_at
		FROM post_comment
		WHERE post_uuid = $1
	`

	return r.tx.Get(ctx, comment, query, postUUID)
}

func (r *comntRepo) GetCommentByUUID(ctx context.Context, comment *entity.Comment, commentUUID uuid.UUID) error {
	query := `
		SELECT
			uuid,
			user_uuid,
			post_uuid,
			comment,
			created_at,
			deleted_at
		FROM post_comment
		WHERE uuid = $1
	`

	return r.tx.GetOne(ctx, comment, query, commentUUID)
}

func (r *comntRepo) DeleteComment(ctx context.Context, commentUUID uuid.UUID) error {
	query := `
		DELETE FROM post_comment
		WHERE uuid = $1
	`

	return r.tx.Delete(ctx, query, commentUUID)
}
