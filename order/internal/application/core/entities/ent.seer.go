package entity

import "github.com/google/uuid"

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
