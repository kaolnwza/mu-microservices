package entity

import "github.com/google/uuid"

type UserWallet struct {
	UserUUID uuid.UUID `db:"user_uuid" json:"user_uuid"`
	Fund     int64     `db:"fund" json:"fund"`
}
