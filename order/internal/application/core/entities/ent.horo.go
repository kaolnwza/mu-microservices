package entity

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type HoroLocation string

const (
	HORO_LOCATION_HYBIRD HoroLocation = "hybrid"
	HORO_LCOATION_ONSITE HoroLocation = "onsite"
	HORO_LCOATION_ONLINE HoroLocation = "online"
)

var HoroLocationMap = map[string]HoroLocation{
	"hybrid": HORO_LOCATION_HYBIRD,
	"onsite": HORO_LCOATION_ONSITE,
	"online": HORO_LCOATION_ONLINE,
}

type HoroType string

const (
	HORO_TYPE_GRPSY_CARD  HoroType = "gypsy_card"
	HORO_TYPE_ASTROLOGY   HoroType = "astrology"
	HORO_TYPE_PHYSIOGNOMY HoroType = "physiognomy"
	HORO_TYPE_TARO_CARD   HoroType = "taro_card"
)

var HoroTypeMap = map[string]HoroType{
	"gypsy_card":  HORO_TYPE_GRPSY_CARD,
	"astrology":   HORO_TYPE_ASTROLOGY,
	"physiognomy": HORO_TYPE_PHYSIOGNOMY,
	"taro_card":   HORO_TYPE_TARO_CARD,
}

type Horo struct {
	UUID            uuid.UUID      `db:"uuid"`
	SeerUUID        uuid.UUID      `db:"seer_uuid"`
	HoroLocation    HoroLocation   `db:"horo_location"`
	HoroType        HoroType       `db:"horo_type"`
	Title           sql.NullString `db:"title"`
	Description     sql.NullString `db:"description"`
	Price           int            `db:"price"`
	Available       bool           `db:"available"`
	MeetingStatus   bool           `db:"meeting_status"`
	MeetingLocation sql.NullString `db:"meeting_location"`
	ChatStatus      bool           `db:"chat_status"`
	VoiceCallStatus bool           `db:"voice_call_status"`
	VideoCallStatus bool           `db:"video_call_status"`
}

type HoroSchedule struct {
	UUID            uuid.UUID `db:"uuid" json:"uuid"`
	HoroServiceUUID uuid.UUID `db:"horo_service_uuid" json:"horo_service_uuid"`
	StartTime       time.Time `db:"start_time" json:"start_time"`
	EndTime         time.Time `db:"end_time" json:"end_time"`
}

type HoroScheduleTime struct {
	StartTime time.Time `db:"start_time" json:"start_time"`
	EndTime   time.Time `db:"end_time" json:"end_time"`
}

type HoroAvailableEvent struct {
	HoroLocation    HoroLocation        `db:"horo_location"`
	MeetingStatus   bool                `db:"meeting_status"`
	MeetingLocation string              `db:"meeting_location"`
	ChatStatus      bool                `db:"chat_status"`
	VoiceCallStatus bool                `db:"voice_call_status"`
	VideoCallStatus bool                `db:"video_call_status"`
	Schedule        []*HoroScheduleTime `db:"schedule" json:"schedule"`
}

type HoroResponse struct {
	UUID            uuid.UUID     `json:"uuid"`
	SeerUUID        uuid.UUID     `json:"seer_uuid"`
	HoroLocation    HoroLocation  `json:"horo_location"`
	HoroType        HoroType      `json:"horo_type"`
	Title           string        `json:"title"`
	Description     string        `json:"description"`
	Price           int           `json:"price"`
	Available       bool          `json:"available"`
	MeetingStatus   bool          `json:"meeting_status"`
	MeetingLocation LineString    `json:"meeting_location"`
	ChatStatus      bool          `json:"chat_status"`
	VoiceCallStatus bool          `json:"voice_call_status"`
	VideoCallStatus bool          `json:"video_call_status"`
	Images          []*HoroImages `json:"horo_images"`
}

type HoroImagesRequest struct {
	UploadUUID uuid.UUID `json:"upload_uuid"`
	ImageOrder int32     `json:"image_order"`
}

type HoroImages struct {
	Url        string `json:"image_url"`
	ImageOrder int32  `json:"image_order"`
}
