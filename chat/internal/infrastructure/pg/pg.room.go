package pg

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/chat/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/chat/internal/ports"
)

type roomRepo struct {
	tx port.Transactor
}

func NewRoomRepository(tx port.Transactor) port.RoomRepository {
	return &roomRepo{tx: tx}
}

func (r *roomRepo) CreateRoom(ctx context.Context, stUserUUID uuid.UUID, ndUserUUID uuid.UUID, startTime string, endTime string) error {
	query := `
		WITH room_created AS (
			INSERT INTO room (start_time, end_time)	
			VALUES ($1, $2)
			RETURNING uuid
		)

		INSERT INTO room_user (room_uuid, user_uuid)
		VALUES 
			((SELECT uuid FROM room_created), $3),
			((SELECT uuid FORM room_created), $4)
	`

	return r.tx.Insert(ctx, query, startTime, endTime, stUserUUID, ndUserUUID)
}

func (r *roomRepo) GetRoomByUserUUID(ctx context.Context, dest *[]*entity.Room, userUUID uuid.UUID) error {
	query := `
		SELECT
			room.uuid,
			start_time,
			end_time,
			user_uuid AS "meeter_uuid",
			created_at
		FROM room
		LEFT JOIN room_user ON room_uuid = room.uuid AND user_uuid != $1
		WHERE EXISTS (
			SELECT 1 FROM room_user
			WHERE user_uuid = $1
			AND room_uuid = room.uuid
		)
		AND now() < start_time
		AND status IS TRUE
		ORDER BY created_at DESC
	`

	return r.tx.Get(ctx, dest, query, userUUID)
}

func (r *roomRepo) GetRoomByUUID(ctx context.Context, dest *entity.Room, userUUID uuid.UUID, roomUUID uuid.UUID) error {
	query := `
		SELECT
			room.uuid,
			start_time,
			end_time,
			user_uuid AS "meeter_uuid",
			created_at,
			status
		FROM room
		LEFT JOIN room_user ON room_uuid = room.uuid AND user_uuid != $1
		WHERE EXISTS (
			SELECT 1 FROM room_user
			WHERE user_uuid = $1
			AND room_uuid = room.uuid
		)
		AND room.uuid = $2
	`

	return r.tx.GetOne(ctx, dest, query, userUUID, roomUUID)
}

func (r *roomRepo) UpdateRoomStatusByOrderUUID(ctx context.Context, orderUUID uuid.UUID, status bool) error {
	query := `
		UPDATE room
		SET 
			status = $1
		WHERE uuid = $2
	`

	return r.tx.Update(ctx, query, status, orderUUID)
}
