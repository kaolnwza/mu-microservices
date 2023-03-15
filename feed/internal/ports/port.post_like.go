package port

import (
	"context"

	"github.com/google/uuid"
)

type LikeRepository interface {
	PostLike(ctx context.Context, status *bool, userUUID uuid.UUID, postUUID uuid.UUID) error
	PostUnlike(ctx context.Context, status *bool, userUUID uuid.UUID, postUUID uuid.UUID) error
}

type LikeService interface {
	PostLike(ctx context.Context, userUUID uuid.UUID, postUUID uuid.UUID) (*bool, error)
	PostUnlike(ctx context.Context, userUUID uuid.UUID, postUUID uuid.UUID) (*bool, error)
}
