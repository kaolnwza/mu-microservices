package entity

import (
	"time"

	"github.com/google/uuid"
)

type RoomMessage struct {
	UUID      uuid.UUID `db:"uuid" json:"uuid"`
	RoomUUID  uuid.UUID `db:"room_uuid" json:"room_uuid"`
	UserUUID  uuid.UUID `db:"user_uuid" json:"user_uuid"`
	Message   string    `db:"message" json:"message"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type RoomMessageResponse struct {
	UUID      uuid.UUID `json:"uuid"`
	RoomUUID  uuid.UUID `json:"room_uuid"`
	UserUUID  uuid.UUID `json:"user_uuid"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
