package pg

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/chat/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/chat/internal/ports"
)

type roomMsgRepo struct {
	tx port.Transactor
}

func NewRoomMessageRepository(tx port.Transactor) port.RoomMessageRepository {
	return &roomMsgRepo{tx: tx}
}

func (r *roomMsgRepo) CreateRoomMessage(ctx context.Context, roomUUID uuid.UUID, userUUID uuid.UUID, message string) error {
	query := `
		INSERT INTO room_message (room_uuid, user_uuid, message)
		VALUES ($1, $2, $3)
	`

	return r.tx.Insert(ctx, query, roomUUID, userUUID, message)
}

func (r *roomMsgRepo) GetMessageByRoomUUID(ctx context.Context, msg *[]*entity.RoomMessage, roomUUID uuid.UUID, timeOffset string) error {
	query := `
		SELECT 
			uuid,
			room_uuid,
			user_uuid,
			message,
			created_at
		FROM room_message
		WHERE room_uuid = $1
		AND created_at < $2
		ORDER BY created_at DESC
		LIMIT 20
	`

	return r.tx.Get(ctx, msg, query, roomUUID, timeOffset)

}
