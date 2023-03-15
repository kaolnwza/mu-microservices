package entity

import "github.com/google/uuid"

type UserResponse struct {
	UUID              uuid.UUID `json:"uuid"`
	DisplayName       string    `json:"display_name"`
	Birthday          string    `json:"birthday"`
	Description       string    `json:"description"`
	TelNumber         string    `json:"tel_number"`
	Email             string    `json:"email"`
	Role              string    `json:"role"`
	ProfilePictureURL string    `json:"profile_picture_url"`
	// CoinAmount        int            `db:"coin_amount" json:"coin_amount"`
	// Upload *Upload `db:"profile_picture" json:"profile_picture"`
}
