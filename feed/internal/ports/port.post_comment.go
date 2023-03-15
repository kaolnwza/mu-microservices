package port

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/feed/internal/application/core/entities"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, comment *entity.Comment, userUUID uuid.UUID, postUUID uuid.UUID, text string) error
	DeleteComment(ctx context.Context, commentUUID uuid.UUID) error
	GetCommentByPostUUID(ctx context.Context, comment *[]*entity.Comment, postUUID uuid.UUID) error
	GetCommentByUUID(ctx context.Context, comment *entity.Comment, commentUUID uuid.UUID) error
}

type CommentService interface {
	CreateComment(ctx context.Context, userUUID uuid.UUID, postUUID uuid.UUID, text string) (*entity.CommentResponse, error)
	DeleteCommentByUUID(ctx context.Context, userUUID uuid.UUID, commentUUID uuid.UUID) error
	GetCommentByPostUUID(ctx context.Context, postUUID uuid.UUID) (*[]*entity.CommentResponse, error)
}
