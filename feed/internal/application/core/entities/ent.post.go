package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	UUID          uuid.UUID    `db:"uuid"`
	UserUUID      uuid.UUID    `db:"user_uuid"`
	Title         string       `db:"title"`
	Text          string       `db:"text"`
	CreatedAt     time.Time    `db:"created_at"`
	DeletedAt     sql.NullTime `db:"deleted_at"`
	LikeStatus    bool         `db:"like_status"`
	LikeCount     int          `db:"like_count"`
	CommentAmount int          `db:"comment_amount"`
}

type PostImage struct {
	UUID       uuid.UUID     `db:"uuid"`
	PostUUID   uuid.UUID     `db:"post_uuid"`
	UploadUUID uuid.NullUUID `db:"upload_uuid" json:"upload_uuid"`
	Order      sql.NullInt32 `db:"order" json:"order"`
	CreatedAt  time.Time     `db:"created_at"`
}

type PostWithImage struct {
	*Post      `db:"post"`
	*PostImage `db:"post_image"`
}

type PostImageResponse struct {
	UUID       uuid.UUID `json:"uuid"`
	Order      int       `json:"order"`
	UploadUUID uuid.UUID `json:"-"`
	Url        *string   `json:"url"`
}

type PostImageRequest struct {
	UploadUUID uuid.UUID `json:"upload_uuid"`
	Order      int       `json:"order"`
}

type PostResponse struct {
	UUID          uuid.UUID            `json:"uuid"`
	UserUUID      uuid.UUID            `json:"user_uuid"`
	Title         string               `json:"title"`
	Text          string               `json:"text"`
	CreatedAt     time.Time            `json:"created_at"`
	LikeStatus    bool                 `json:"like_status"`
	LikeCount     int                  `json:"like_count"`
	CommentAmount int                  `json:"comment_amount"`
	PostImage     []*PostImageResponse `json:"post_images"`
}
