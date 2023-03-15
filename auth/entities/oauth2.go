package entity

import "github.com/google/uuid"

type OAuth2 struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Locale        string `json:"locale"`
}

type AuthGoogle struct {
	UUID     uuid.UUID `db:"uuid" json:"uuid"`
	UserUUID uuid.UUID `db:"user_uuid" json:"user_uuid"`
	Email    string    `db:"email" json:"email"`
}
