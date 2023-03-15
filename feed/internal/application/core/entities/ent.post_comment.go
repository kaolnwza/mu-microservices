package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	UUID      uuid.UUID    `db:"uuid" json:"uuid"`
	UserUUID  uuid.UUID    `db:"user_uuid" json:"user_uuid"`
	PostUUID  uuid.UUID    `db:"post_uuid" json:"post_uuid"`
	Comment   string       `db:"comment" json:"comment"`
	CreatedAt time.Time    `db:"created_at" json:"created_at"`
	DeletedAt sql.NullTime `db:"deleted_at" json:"deleted_at"`
}

type CommentResponse struct {
	UUID      uuid.UUID `json:"uuid"`
	UserUUID  uuid.UUID `json:"user_uuid"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
