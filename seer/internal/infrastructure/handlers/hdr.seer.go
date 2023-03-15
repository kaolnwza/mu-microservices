package handler

import (
	"net/http"

	"github.com/google/uuid"
	port "github.com/kaolnwza/muniverse/seer/internal/ports"
	log "github.com/kaolnwza/muniverse/seer/lib/logs"
)

type seerHandler struct {
	seerSvc port.SeerService
}

func NewSeerHandler(svc port.SeerService) seerHandler {
	return seerHandler{seerSvc: svc}
}

func (h *seerHandler) GetSeerByUUID(c port.Context) {
	seerUUID, err := uuid.Parse(c.Param("seer_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	seer, err := h.seerSvc.GetSeerByUserUUID(c.Ctx(), seerUUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &seer)
}
