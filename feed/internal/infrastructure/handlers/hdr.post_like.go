package handler

import (
	"net/http"

	"github.com/google/uuid"
	port "github.com/kaolnwza/muniverse/feed/internal/ports"
	log "github.com/kaolnwza/muniverse/feed/lib/logs"
)

type likeHdr struct {
	likeSvc port.LikeService
}

func NewLikeHandler(likeSvc port.LikeService) *likeHdr {
	return &likeHdr{likeSvc: likeSvc}
}

func (h *likeHdr) PostLikeHandler(c port.Context) {
	userUUID := c.AccessUserUUID()
	postUUID, err := uuid.Parse(c.Param("post_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	status, err := h.likeSvc.PostLike(c.Ctx(), userUUID, postUUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"status": status,
	})
}

func (h *likeHdr) PostUnlikeHandler(c port.Context) {
	userUUID := c.AccessUserUUID()
	postUUID, err := uuid.Parse(c.Param("post_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	status, err := h.likeSvc.PostUnlike(c.Ctx(), userUUID, postUUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": status,
	})
}
