package entity

import "github.com/google/uuid"

type RoomUser struct {
	UUID     uuid.UUID `db:"uuid"`
	RoomUUID uuid.UUID `db:"room_uuid"`
	UserUUID uuid.UUID `db:"user_uuid"`
}
