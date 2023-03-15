package handler

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	port "github.com/kaolnwza/muniverse/order/internal/ports"
	log "github.com/kaolnwza/muniverse/order/lib/logs"
)

type horoHdr struct {
	svc port.HoroService
}

func NewHoroHandler(svc port.HoroService) horoHdr {
	return horoHdr{svc: svc}
}

func (h *horoHdr) CreateHoroServiceHandler(c port.Context) {
	userUUID := c.AccessUserUUID()
	title := c.FormValue("title")
	desc := c.FormValue("description")
	latitude := c.FormValue("latitude")
	longtitude := c.FormValue("longtitude")

	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	horoLocation := c.FormValue("horo_location")
	isMeet, err := strconv.ParseBool(c.FormValue("meeting_status"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	isChat, err := strconv.ParseBool(c.FormValue("chat_status"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	isVideoCall, err := strconv.ParseBool(c.FormValue("video_call_status"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	isVoiceCall, err := strconv.ParseBool(c.FormValue("voice_call_status"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	imagesJSON := c.FormValue("images")

	if err := h.svc.CreateHoroService(c.Ctx(), userUUID, title, desc, price, horoLocation, isMeet, latitude, longtitude, isChat, isVoiceCall, isVideoCall, imagesJSON); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (h *horoHdr) UpdateHoroServiceHandler(c port.Context) {
	horoUUID, err := uuid.Parse(c.Param("horo_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	title := c.FormValue("title")
	desc := c.FormValue("description")
	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	horoLocation := c.FormValue("horo_location")
	isMeet, err := strconv.ParseBool(c.FormValue("meeting_status"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	latitude := c.FormValue("latitude")
	longtitude := c.FormValue("longtitude")

	isChat, err := strconv.ParseBool(c.FormValue("chat_status"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	isVideoCall, err := strconv.ParseBool(c.FormValue("video_call_status"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	isVoiceCall, err := strconv.ParseBool(c.FormValue("voice_call_status"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.svc.UpdateHoroService(c.Ctx(), horoUUID, title, desc, price, horoLocation, isMeet, latitude, longtitude, isChat, isVoiceCall, isVideoCall); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *horoHdr) UpdateHoroOnEventHandler(c port.Context) {
	horoUUID, err := uuid.Parse(c.Param("horo_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	title := c.FormValue("title")
	desc := c.FormValue("description")
	horoLocation := c.FormValue("horo_location")
	eventTime := c.FormValue("event_time")

	if err := h.svc.UpdateHoroServiceOnEvent(c.Ctx(), horoUUID, title, desc, horoLocation, eventTime); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *horoHdr) UpdateHoroStatusHandler(c port.Context) {
	horoUUID, err := uuid.Parse(c.Param("horo_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	status, err := strconv.ParseBool(c.FormValue("available"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.svc.UpdateHoroServiceStatus(c.Ctx(), horoUUID, status); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *horoHdr) GetHoroAvailableEventByDateHandler(c port.Context) {
	horoUUID, err := uuid.Parse(c.Param("horo_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	date := c.FormValue("date")

	resp, err := h.svc.GetHoroAvailableEventByDate(c.Ctx(), horoUUID, date)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *horoHdr) GetHoroScheduleEventByDate(c port.Context) {
	horoUUID, err := uuid.Parse(c.Param("horo_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	date := c.FormValue("date")

	resp, err := h.svc.GetHoroScheduleEventByDate(c.Ctx(), horoUUID, date)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *horoHdr) GetHoroByHoroUUIDHandler(c port.Context) {
	horoUUID, err := uuid.Parse(c.Param("horo_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.svc.GetHoroByHoroUUID(c.Ctx(), horoUUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *horoHdr) GetAllHoroServiceHandler(c port.Context) {
	resp, err := h.svc.GetAllHoroService(c.Ctx())
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *horoHdr) GetAllHoroServiceBySeerUUIDHandler(c port.Context) {
	seerUUID, err := uuid.Parse(c.Param("seer_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.svc.GetAllHoroServiceBySeerUUID(c.Ctx(), seerUUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
