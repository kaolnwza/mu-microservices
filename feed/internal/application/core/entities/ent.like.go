package entity

import "github.com/google/uuid"

type PostLike struct {
	UUID     uuid.UUID `db:"uuid"`
	PostUUID uuid.UUID `db:"post_uuid"`
}
