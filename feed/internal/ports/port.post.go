package port

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/feed/internal/application/core/entities"
)

type PostRepository interface {
	GetByTime(ctx context.Context, post *[]*entity.PostWithImage, bottomTime string, userUUID uuid.UUID) error
	GetByPostUUID(ctx context.Context, post *[]*entity.PostWithImage, postUUID uuid.UUID, userUUID uuid.UUID) error
	CreatePost(ctx context.Context, postUUID *uuid.UUID, title string, text string, userUUID uuid.UUID) error
	CreatePostImages(ctx context.Context, postUUID *uuid.UUID, images []*entity.PostImageRequest) error
	DeletePost(ctx context.Context, postUUID uuid.UUID) error
	UpdatePost(ctx context.Context, postUUID uuid.UUID, title string, desc string) error
}

type PostService interface {
	GetAllPosts(ctx context.Context, bottomTime string, userUUID uuid.UUID) (*[]*entity.PostResponse, error)
	GetPostByPostUUID(ctx context.Context, postUUID uuid.UUID, userUUID uuid.UUID) (*entity.PostResponse, error)
	CreatePost(ctx context.Context, title string, text string, userUUID uuid.UUID, images []*entity.PostImageRequest) error
	DeletePost(ctx context.Context, postUUID uuid.UUID, userUUID uuid.UUID) error
	UpdatePost(ctx context.Context, userUUID uuid.UUID, postUUID uuid.UUID, title string, desc string) error
}
