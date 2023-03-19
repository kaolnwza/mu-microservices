package handler

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	port "github.com/kaolnwza/muniverse/chat/internal/ports"
	log "github.com/kaolnwza/muniverse/chat/lib/logs"
)

type roomMsgHdr struct {
	roomMsgSvc port.RoomMessageService
}

func NewRoomMessageHandler(roomMsgSvc port.RoomMessageService) *roomMsgHdr {
	return &roomMsgHdr{roomMsgSvc: roomMsgSvc}
}

func (h *roomMsgHdr) GetRoomMessageByRoomUUIDHandler(c port.Context) {
	roomUUID, err := uuid.Parse(c.Param("room_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	timeOffset := c.Request().URL.Query().Get("time_offset")
	if timeOffset == "" {
		timeOffset = time.Now().Format(time.RFC3339)
	}

	roomMsg, err := h.roomMsgSvc.GetMessageByRoomUUID(c.Ctx(), roomUUID, timeOffset)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &roomMsg)
}
