package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/feed/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/feed/internal/ports"
)

type postRepo struct {
	tx port.Transactor
}

func NewPostRepository(tx port.Transactor) port.PostRepository {
	return &postRepo{tx: tx}
}

func (r *postRepo) GetByTime(ctx context.Context, post *[]*entity.PostWithImage, bottomTime string, userUUID uuid.UUID) error {
	query := `
		WITH posts AS (
			SELECT 
				uuid,
				user_uuid,
				title,
				text,
				created_at,
				deleted_at,
				COALESCE((SELECT TRUE FROM post_like WHERE user_uuid = $1), FALSE) like_status,
				COALESCE((SELECT COUNT(1) FROM post_like WHERE post_uuid = post.uuid AND post_uuid = post.uuid), 0) like_count,
				COALESCE((SELECT COUNT(1) FROM post_comment WHERE post_uuid = post.uuid), 0) comment_amount
			FROM post 
			WHERE deleted_at IS NULL`
	if bottomTime != "" {
		query += `
			AND created_at < '` + bottomTime + `' `
	}
	query += `
			LIMIT 10
		)

		SELECT
			posts.uuid AS "post.uuid",
			user_uuid AS "post.user_uuid",
			title AS "post.title",
			text AS "post.text",
			posts.created_at AS "post.created_at",
			posts.deleted_at AS "post.deleted_at",
			like_status AS "post.like_status",
			like_count AS "post.like_count",
			comment_amount AS "post.comment_amount",
			upload_uuid AS "post_image.upload_uuid",
			post_image.order AS "post_image.order",
			post_image.uuid AS "post_image.uuid"
		FROM posts
		LEFT JOIN post_image ON post_image.post_uuid = posts.uuid AND post_image.deleted_at IS NULL
		ORDER BY posts.created_at DESC, posts.uuid, post_image.order
	`

	return r.tx.Get(ctx, post, query, userUUID)
}

func (r *postRepo) GetByPostUUID(ctx context.Context, post *[]*entity.PostWithImage, postUUID uuid.UUID, userUUID uuid.UUID) error {
	query := `
		WITH posts AS (
			SELECT 
				uuid,
				user_uuid,
				title,
				text,
				created_at,
				deleted_at,
				COALESCE((SELECT TRUE FROM post_like WHERE user_uuid = $2 AND post_uuid = post.uuid), FALSE) like_status,
				COALESCE((SELECT COUNT(1) FROM post_like WHERE post_uuid = post.uuid), 0) like_count,
				COALESCE((SELECT COUNT(1) FROM post_comment WHERE post_uuid = post.uuid), 0) comment_amount
			FROM post
			WHERE uuid = $1
			AND deleted_at IS NULL
			LIMIT 10
		)

		SELECT
			posts.uuid AS "post.uuid",
			user_uuid AS "post.user_uuid",
			title AS "post.title",
			text AS "post.text",
			posts.created_at AS "post.created_at",
			posts.deleted_at AS "post.deleted_at",
			like_status AS "post.like_status",
			like_count AS "post.like_count",
			comment_amount AS "post.comment_amount",
			upload_uuid AS "post_image.upload_uuid",
			post_image.order AS "post_image.order",
			post_image.uuid AS "post_image.uuid"
		FROM posts
		LEFT JOIN post_image ON post_image.post_uuid = posts.uuid AND post_image.deleted_at IS NULL
		ORDER BY posts.created_at DESC, posts.uuid, post_image.order
`

	return r.tx.Get(ctx, post, query, postUUID, userUUID)
}

func (r *postRepo) CreatePost(ctx context.Context, postUUID *uuid.UUID, title string, text string, userUUID uuid.UUID) error {
	query := `
		INSERT INTO post (title, text, user_uuid)
		VALUES ($1, $2, $3)
		RETURNING uuid
	`

	return r.tx.InsertWithReturningOne(ctx, postUUID, query, title, text, userUUID)
}

func (r *postRepo) CreatePostImages(ctx context.Context, postUUID *uuid.UUID, images []*entity.PostImageRequest) error {
	values := ``

	for idx, item := range images {
		values += fmt.Sprintf(`('%v', '%s', %v)`, postUUID, item.UploadUUID, item.Order)
		if idx != len(images)-1 {
			values += ","
		}
	}

	query := `
		INSERT INTO post_image (post_uuid, upload_uuid, "order")
		VALUES ` + values

	return r.tx.Insert(ctx, query)
}

func (r *postRepo) DeletePost(ctx context.Context, postUUID uuid.UUID) error {
	query := `
		UPDATE post
		SET deleted_at = now()
		WHERE uuid = $1
	`

	return r.tx.Update(ctx, query, postUUID)
}

func (r *postRepo) UpdatePost(ctx context.Context, postUUID uuid.UUID, title string, desc string) error {
	query := `
		UPDATE post
		SET 
			title = $2,
			text = $3

		WHERE uuid = $1
	`

	return r.tx.Update(ctx, query, postUUID, title, desc)
}
