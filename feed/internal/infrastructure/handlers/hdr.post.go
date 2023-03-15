package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/feed/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/feed/internal/ports"
	log "github.com/kaolnwza/muniverse/feed/lib/logs"
)

type postHdr struct {
	postSvc port.PostService
}

func NewPostHandler(postSvc port.PostService) *postHdr {
	return &postHdr{postSvc: postSvc}
}

func (h *postHdr) GetPostByPostUUIDHandler(c port.Context) {
	userUUID := c.AccessUserUUID()
	postUUID, err := uuid.Parse(c.Param("post_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.postSvc.GetPostByPostUUID(c.Ctx(), postUUID, userUUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &resp)
}

func (h *postHdr) GetAllPostHandler(c port.Context) {
	userUUID := c.AccessUserUUID()
	bottomTime := c.FormValue("bottom_time")

	resp, err := h.postSvc.GetAllPosts(c.Ctx(), bottomTime, userUUID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &resp)
}

func (h *postHdr) CreatePostHandler(c port.Context) {
	userUUID := c.AccessUserUUID()
	title := c.FormValue("title")
	text := c.FormValue("text")
	imagesString := c.FormValue("images")

	images := make([]*entity.PostImageRequest, 0)
	if err := json.Unmarshal([]byte(imagesString), &images); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := h.postSvc.CreatePost(c.Ctx(), title, text, userUUID, images); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (h *postHdr) DeletePostHandler(c port.Context) {
	userUUID := c.AccessUserUUID()
	postUUID, err := uuid.Parse(c.Param("post_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.postSvc.DeletePost(c.Ctx(), postUUID, userUUID); err != nil {
		log.Error(err)
		if err.Error() == "not owner kub" {
			//handler on bussiness later
			c.JSON(http.StatusUnauthorized, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *postHdr) UpdatePostHandler(c port.Context) {
	userUUID := c.AccessUserUUID()
	postUUID, err := uuid.Parse(c.Param("post_uuid"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	title := c.FormValue("title")
	text := c.FormValue("text")

	if err := h.postSvc.UpdatePost(c.Ctx(), userUUID, postUUID, title, text); err != nil {
		log.Error(err)
		if err.Error() == "not owner kub" {
			//handler on bussiness later
			c.JSON(http.StatusUnauthorized, err.Error())
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
