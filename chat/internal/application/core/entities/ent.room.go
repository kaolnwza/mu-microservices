package entity

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	UUID          uuid.UUID `db:"uuid" json:"uuid"`
	HoroOrderUUID uuid.UUID `db:"horo_order_uuid" json:"-"`
	StartTime     time.Time `db:"start_time" json:"start_time"`
	EndTime       time.Time `db:"end_time" json:"end_time"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	MeeterUUID    uuid.UUID `db:"meeter_uuid" json:"meeter_uuid"`
	Status        bool      `db:"status" json:"status"`
	LastMessage   string    `db:"last_message" json:"last_message"`
}

type RoomResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	CreatedAt   time.Time `json:"created_at"`
	MeeterUUID  uuid.UUID `json:"meeter_uuid"`
	Status      bool      `json:"status"`
	LastMessage string    `json:"last_message"`
}
