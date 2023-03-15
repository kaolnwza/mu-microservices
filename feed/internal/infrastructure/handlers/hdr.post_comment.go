package handler

import (
	"net/http"

	"github.com/google/uuid"
	port "github.com/kaolnwza/muniverse/feed/internal/ports"
	log "github.com/kaolnwza/muniverse/feed/lib/logs"
)

type comntHdr struct {
	comntSvc port.CommentService
}

func NewCommentHandler(comntSvc port.CommentService) *comntHdr {
	return &comntHdr{comntSvc: comntSvc}
}

func (h *comntHdr) GetCommentByPostUUIDHandler(c port.Context) {
	postUUID, err := uuid.Parse(c.Param("post_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.comntSvc.GetCommentByPostUUID(c.Ctx(), postUUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &resp)
}

func (h *comntHdr) CreateCommentHandler(c port.Context) {
	postUUID, err := uuid.Parse(c.Param("post_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	comment := c.FormValue("comment")
	userUUID := c.AccessUserUUID()

	resp, err := h.comntSvc.CreateComment(c.Ctx(), userUUID, postUUID, comment)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, &resp)
}

func (h *comntHdr) DeleteCommentByUUIDHandler(c port.Context) {
	userUUID := c.AccessUserUUID()
	commentUUID, err := uuid.Parse(c.Param("comment_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.comntSvc.DeleteCommentByUUID(c.Ctx(), userUUID, commentUUID); err != nil {
		log.Error(err)
		if err.Error() == "not owner kub" {
			c.JSON(http.StatusUnauthorized, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
