package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	port "github.com/kaolnwza/muniverse/chat/internal/ports"
	log "github.com/kaolnwza/muniverse/chat/lib/logs"
)

type roomHdr struct {
	wsRoom  port.WSRoom
	roomSvc port.RoomService
}

func NewRoomHandler(wsRoom port.WSRoom, roomSvc port.RoomService) *roomHdr {
	return &roomHdr{wsRoom: wsRoom, roomSvc: roomSvc}
}

func (h *roomHdr) JoinRoomHandler(c port.Context) {
	userUUID := c.AccessUserUUID()
	roomUUIDs := c.Param("room_uuid")
	roomUUID, err := uuid.Parse(roomUUIDs)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	room, err := h.roomSvc.GetRoomByUUID(c.Ctx(), userUUID, roomUUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if !(room.StartTime.Before(time.Now()) && room.EndTime.After(time.Now())) {
		log.Error(errors.New("out of meeting time"))
		c.JSON(http.StatusBadRequest, map[string]string{
			"msg": "out of meeting time, connect without websocket!!",
		})
		return
	}

	wsConn, err := h.wsRoom.RoomUpgrader(c.Writer(), c.Request())
	if err != nil {
		log.Error(err)
		c.JSON(500, err.Error())
		return
	}

	h.wsRoom.ActiveRoom(roomUUIDs)

	// h.wsRoom.JoinRoom(wsConn, "asd", qu.Get("user_uuid"))
	h.wsRoom.JoinRoom(wsConn, roomUUIDs, userUUID.String())

}

func (h *roomHdr) GetRoomByUserUUIDHandler(c port.Context) {
	userUUID := c.AccessUserUUID()
	rooms, err := h.roomSvc.GetRoomByUserUUID(c.Ctx(), userUUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, rooms)
}
