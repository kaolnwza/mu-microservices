package handler

import (
	"net/http"

	"github.com/google/uuid"
	port "github.com/kaolnwza/muniverse/user/internal/ports"
	log "github.com/kaolnwza/muniverse/user/lib/logs"
)

type userHdr struct {
	userSvc port.UserService
}

func NewUserHandler(s port.UserService) userHdr {
	return userHdr{userSvc: s}
}

func (h *userHdr) Health(c port.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"Hello": "World",
	})
}

func (h *userHdr) GetUserByUUIDHandler(c port.Context) {
	userUUID, err := uuid.Parse(c.Param("user_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.userSvc.GetUserByUUID(c.Ctx(), userUUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)

}

func (h *userHdr) GetUserByTokenHandler(c port.Context) {
	userUUID := c.AccessUserUUID()

	user, err := h.userSvc.GetUserByUUID(c.Ctx(), userUUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)

}

func (h *userHdr) UpdateUserHandler(c port.Context) {
	userUUID := c.AccessUserUUID()
	displayName := c.Request().FormValue("display_name")
	dob := c.Request().FormValue("birthday")
	desc := c.Request().FormValue("description")
	uploadUUID := uuid.Nil
	if uploadUUIDstr := c.Request().FormValue("upload_uuid"); uploadUUIDstr != "" {
		uploadUUID, _ = uuid.Parse(uploadUUIDstr)
	}

	if err := h.userSvc.UpdateUser(c.Ctx(), userUUID, displayName, dob, desc, uploadUUID); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
