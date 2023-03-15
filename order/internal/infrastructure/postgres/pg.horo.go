package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/order/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/order/internal/ports"
)

type horoRepo struct {
	tx port.Transactor
}

func NewHoroRepository(tx port.Transactor) port.HoroRepository {
	return &horoRepo{tx: tx}
}

func (r *horoRepo) Create(
	ctx context.Context,
	seerUUID uuid.UUID,
	title string,
	desc string,
	price int,
	horoLocation entity.HoroLocation,
	isMeet bool,
	meetLocation string,
	isChat bool,
	isVoiceCall bool,
	isVideoCall bool,
	dest *entity.Horo) error {

	query := `
	INSERT INTO horo_service (
		seer_uuid, title, description, price, horo_location, meeting_status, meeting_location, chat_status, voice_call_status, video_call_status
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING uuid
	`

	return r.tx.InsertWithReturningOne(ctx, dest, query, seerUUID, title, desc, price, horoLocation, isMeet, meetLocation, isChat, isVoiceCall, isVideoCall)
}

func (r *horoRepo) CreateSchedule(ctx context.Context, sche []entity.HoroSchedule, horoUUID uuid.UUID) error {
	query := `
		INSERT INTO horo_service_schedule (horo_service_uuid, start_time, end_time)
		VALUES 
	`
	for idx, item := range sche {
		query += fmt.Sprintf(`('%s', '%s', '%s')`, horoUUID, item.StartTime.Format(time.RFC3339), item.EndTime.Format(time.RFC3339))
		// query += `('` + horoUUID.String() + `', ` + `'` + (item.StartTime).Format(time.RFC3339) + `',` + `'` + (item.StartTime).Format(time.RFC3339) + `')`
		if idx != len(sche)-1 {
			query += ","
		}
	}

	return r.tx.Insert(ctx, query)
}

func (r *horoRepo) Update(
	ctx context.Context,
	horoUUID uuid.UUID,
	title string,
	desc string,
	price int,
	horoLocation entity.HoroLocation,
	isMeet bool,
	meetLocation string,
	isChat bool,
	isVoiceCall bool,
	isVideoCall bool) error {
	query := `
		UPDATE horo_service
		SET
			title = $1, 
			description = $2, 
			price = $3, 
			horo_location = $4, 
			meeting_status = $5, 
			meeting_location = $6, 
			chat_status = $7, 
			voice_call_status = $8, 
			video_call_status = $9
		WHERE uuid = $10
	`

	return r.tx.Update(ctx, query, title, desc, price, horoLocation, isMeet, meetLocation, isChat, isVoiceCall, isVideoCall, horoUUID)
}

func (r *horoRepo) UpdateOnEvent(ctx context.Context, horoUUID uuid.UUID, title string, desc string, horoLocation entity.HoroLocation) error {
	query := `
		UPDATE horo_service
		SET 
			title = $2,
			description = $3,
			horo_location = $4
		WHERE uuid = $1
	`

	return r.tx.Update(ctx, query, horoUUID, title, desc, horoLocation)
}

func (r *horoRepo) UpdateStatus(ctx context.Context, horoUUID uuid.UUID, available bool) error {
	query := `
		UPDATE horo_service
		SET 
			available = $2
		WHERE uuid = $1
	`

	return r.tx.Insert(ctx, query, horoUUID, available)
}

func (r *horoRepo) GetAvailableEventByDate(ctx context.Context, dest *[]*entity.HoroScheduleTime, horoUUID uuid.UUID, date string) error {
	query := `
	WITH sche AS (
		SELECT 
			start_time,
			end_time
		FROM horo_service_schedule
		WHERE horo_service_uuid = $1
		AND start_time > now()
		AND date(start_time) = $2
	),
	free_time AS (
		SELECT 
			generate_series(start_time, end_time - interval '1 hour', INTERVAL '1 hour') start_time
		FROM sche
	)
	
	SELECT 
		start_time,
		start_time + INTERVAL '1 hour' AS end_time
	FROM free_time
	WHERE NOT EXISTS (
		SELECT 1 
		FROM horo_order
		WHERE horo_service_uuid = $1
		AND horo_order.start_time = free_time.start_time
		AND status IN ('fail', 'refund', 'success')
		AND date(start_time) = $2
	)
	`

	return r.tx.Get(ctx, dest, query, horoUUID, date)
}

func (r *horoRepo) GetScheduleEventByDate(ctx context.Context, dest *[]*entity.HoroScheduleTime, horoUUID uuid.UUID, date string) error {
	query := `
		SELECT 
			start_time,
			end_time
		FROM horo_service_schedule
		WHERE horo_service_uuid = $1
		AND date(start_time) = $2
	`

	return r.tx.Get(ctx, dest, query, horoUUID, date)
}

func (r *horoRepo) GetHoroByHoroUUID(ctx context.Context, dest *entity.Horo, horoUUID uuid.UUID) error {
	query := `
		SELECT
			uuid,
			seer_uuid,
			horo_location,
			horo_type,
			title,
			description,
			price,
			available,
			meeting_status,
			ST_AsGeoJSON("meeting_location")::json->>'coordinates' AS "meeting_location",
			chat_status,
			voice_call_status,
			video_call_status
		FROM horo_service
		WHERE horo_service.uuid = $1
	`

	return r.tx.GetOne(ctx, dest, query, horoUUID)
}

func (r *horoRepo) GetAllHoro(ctx context.Context, dest *[]*entity.Horo) error {
	query := `
	SELECT
		uuid,
		seer_uuid,
		horo_location,
		horo_type,
		title,
		description,
		price,
		available,
		meeting_status,
		ST_AsGeoJSON("meeting_location")::json->>'coordinates' AS "meeting_location",
		chat_status,
		voice_call_status,
		video_call_status
	FROM horo_service
`

	return r.tx.Get(ctx, dest, query)
}

func (r *horoRepo) GetAllHoroBySeerUUID(ctx context.Context, dest *[]*entity.Horo, seerUUID uuid.UUID) error {
	query := `
	SELECT
		uuid,
		seer_uuid,
		horo_location,
		horo_type,
		title,
		description,
		price,
		available,
		meeting_status,
		ST_AsGeoJSON("meeting_location")::json->>'coordinates' AS "meeting_location",
		chat_status,
		voice_call_status,
		video_call_status
	FROM horo_service
	WHERE seer_uuid = $1
`

	return r.tx.Get(ctx, dest, query, seerUUID)
}
