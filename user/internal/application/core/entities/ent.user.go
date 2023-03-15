package entity

import (
	"database/sql"

	"github.com/google/uuid"
)

type UserRole string

const (
	USER_ROLE_USER       UserRole = "user"
	USER_ROLE_SIR_VARUTH UserRole = "sir_varuth"
)

type User struct {
	UUID        uuid.UUID      `db:"uuid" json:"uuid"`
	DisplayName string         `db:"display_name" json:"display_name"`
	Birthday    sql.NullTime   `db:"birthday" json:"birthday"`
	Description sql.NullString `db:"description" json:"description"`
	TelNumber   sql.NullString `db:"tel_number" json:"tel_number"`
	Role        UserRole       `db:"role" json:"role"`
	// ProfilePicture    uuid.NullUUID  `db:"profile_picture" json:"-"`
	// ProfilePictureUri sql.NullString `json:"profile_picture_uri"`
	// CoinAmount        int            `db:"coin_amount" json:"coin_amount"`
	// Upload *Upload `db:"profile_picture" json:"profile_picture"`
}

type UserResponse struct {
	UUID              uuid.UUID `json:"uuid"`
	DisplayName       string    `json:"display_name"`
	Birthday          string    `json:"birthday"`
	Description       string    `json:"description"`
	TelNumber         string    `json:"tel_number"`
	Role              UserRole  `json:"role"`
	ProfilePictureURL string    `json:"profile_picture_url"`
	// CoinAmount        int            `db:"coin_amount" json:"coin_amount"`
	// Upload *Upload `db:"profile_picture" json:"profile_picture"`
}
