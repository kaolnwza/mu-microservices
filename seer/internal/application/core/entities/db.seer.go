package entity

import (
	"database/sql"

	"github.com/google/uuid"
)

type Seer struct {
	UUID               uuid.UUID      `db:"uuid"`
	UserUUID           uuid.UUID      `db:"user_uuid"`
	OnsiteAvailable    bool           `db:"onsite_available"`
	ChatAvailable      bool           `db:"chat_available"`
	CallAvailable      bool           `db:"call_available"`
	VideoCallAvailable bool           `db:"video_call_available"`
	Major              string         `db:"major"`
	MajorDescription   sql.NullString `db:"major_description"`
	DescriptionProfile sql.NullString `db:"description_profile"`
	MapCoordinate      sql.NullString `db:"map_coordinate"`
}

type SeerResponse struct {
	UUID               uuid.UUID `json:"uuid"`
	UserUUID           uuid.UUID `json:"user_uuid"`
	OnsiteAvailable    bool      `json:"onsite_available"`
	ChatAvailable      bool      `json:"chat_available"`
	CallAvailable      bool      `json:"call_available"`
	VideoCallAvailable bool      `json:"video_call_available"`
	Major              string    `json:"major"`
	MajorDescription   string    `json:"major_description"`
	DescriptionProfile string    `json:"description_profile"`
	MapCoordinate      string    `json:"map_coordinate"`
	ImageURL           string    `json:"image_url"`
	DisplayName        string    `json:"display_name"`
}
