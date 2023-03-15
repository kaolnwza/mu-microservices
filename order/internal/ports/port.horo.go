package port

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/order/internal/application/core/entities"
)

// 7	price	int4	NO	NULL	NULL		NULL
// 8	available	bool	NO	NULL	true		NULL
// 9	meeting_status	bool	NO	NULL	false		NULL
// 10	meeting_location	varchar(255)	YES	NULL	NULL		NULL
// 11	chat_status	bool	NO	NULL	false		NULL
// 12	voice_call_status	bool	NO	NULL	false		NULL
// 13	video_call_status	bool	NO	NULL	false		NULL
type HoroRepository interface {
	Create(ctx context.Context, userUUID uuid.UUID, title string, desc string, price int, horoLocation entity.HoroLocation, isMeet bool, meetLocation string, isChat bool, isVoiceCall bool, isVideoCall bool, dest *entity.Horo) error
	CreateSchedule(ctx context.Context, sche []entity.HoroSchedule, horoUUID uuid.UUID) error
	Update(ctx context.Context, horoUUID uuid.UUID, title string, desc string, price int, horoLocation entity.HoroLocation, isMeet bool, meetLocation string, isChat bool, isVoiceCall bool, isVideoCall bool) error
	UpdateOnEvent(ctx context.Context, horoUUID uuid.UUID, title string, desc string, horoLocation entity.HoroLocation) error
	UpdateStatus(ctx context.Context, horoUUID uuid.UUID, available bool) error
	GetAvailableEventByDate(ctx context.Context, dest *[]*entity.HoroScheduleTime, horoUUID uuid.UUID, date string) error
	GetScheduleEventByDate(ctx context.Context, dest *[]*entity.HoroScheduleTime, horoUUID uuid.UUID, date string) error
	GetHoroByHoroUUID(ctx context.Context, dest *entity.Horo, horoUUID uuid.UUID) error
	GetAllHoro(ctx context.Context, dest *[]*entity.Horo) error
	GetAllHoroBySeerUUID(ctx context.Context, dest *[]*entity.Horo, seerUUID uuid.UUID) error
}

type HoroService interface {
	CreateHoroService(ctx context.Context, seerUUID uuid.UUID, title string, desc string, price int, horoLocation string, isMeet bool, longitude string, latitude string, isChat bool, isVoiceCall bool, isVideoCall bool, imagesJSON string) error
	UpdateHoroService(ctx context.Context, horoUUID uuid.UUID, title string, desc string, price int, horoLocation string, isMeet bool, longitude string, latitude string, isChat bool, isVoiceCall bool, isVideoCall bool) error
	UpdateHoroServiceOnEvent(ctx context.Context, horoUUID uuid.UUID, title string, desc string, horoLocation string, eventTime string) error
	UpdateHoroServiceStatus(ctx context.Context, horoUUID uuid.UUID, available bool) error
	GetHoroAvailableEventByDate(ctx context.Context, horoUUID uuid.UUID, date string) (*entity.HoroAvailableEvent, error)
	GetHoroScheduleEventByDate(ctx context.Context, horoUUID uuid.UUID, date string) (*[]*entity.HoroScheduleTime, error)
	GetHoroByHoroUUID(ctx context.Context, horoUUID uuid.UUID) (*entity.HoroResponse, error)
	GetAllHoroService(ctx context.Context) (*[]*entity.HoroResponse, error)
	GetAllHoroServiceBySeerUUID(ctx context.Context, seerUUID uuid.UUID) (*[]*entity.HoroResponse, error)
}
