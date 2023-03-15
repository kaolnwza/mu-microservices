package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kaolnwza/muniverse/storage/database"
	log "github.com/kaolnwza/muniverse/storage/logs"
	repository "github.com/kaolnwza/muniverse/storage/repositories"
)

func UploadProfileImg(c *gin.Context) {
	userUUID, ok := c.Request.Context().Value("access_user_uuid").(uuid.UUID)
	if !ok {
		log.Error(fmt.Errorf("not found access_user_uuid"))
		c.Status(http.StatusInternalServerError)
		c.Abort()
		return
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	defer file.Close()

	upload, err := GCSUploadFile(c, userUUID, file, "yesped")
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := repository.UploadImage(database.NewPostgresDB(), &*upload, upload.Path, upload.Bucket, userUUID); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(201, map[string]string{
		"upload_uuid": upload.UUID.String(),
	})

}
func UploadImage(c *gin.Context) {
	userUUID, ok := c.Request.Context().Value("access_user_uuid").(uuid.UUID)
	if !ok {
		log.Error(fmt.Errorf("not found access_user_uuid"))
		c.Status(http.StatusInternalServerError)
		c.Abort()
		return
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	defer file.Close()

	upload, err := GCSUploadFile(c, userUUID, file, "yesped")
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if err := repository.UploadImage(database.NewPostgresDB(), &*upload, upload.Path, upload.Bucket, userUUID); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(201, map[string]string{
		"upload_uuid": upload.UUID.String(),
	})

}
