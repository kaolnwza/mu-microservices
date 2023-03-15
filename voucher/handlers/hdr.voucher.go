package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaolnwza/muniverse/voucher/database"
	entity "github.com/kaolnwza/muniverse/voucher/entities"
	log "github.com/kaolnwza/muniverse/voucher/logs"
	repository "github.com/kaolnwza/muniverse/voucher/repositories"
)

func GetVoucherByCodeHandler(c *gin.Context) {
	code := c.Param("code")
	conn := database.NewPostgresDB()

	voucher := entity.Voucher{}
	if err := repository.GetVoucherByCode(conn, &voucher, code); err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, voucher)
}

// func GetTotalDiscountByCodeHandler(c *gin.Context) {
// 	code := c.Param("code")
// 	conn := database.NewPostgresDB()

// 	voucher := entity.Voucher{}
// 	if err := repository.GetVoucherByCode(conn, &voucher, code); err != nil {
// 		log.Error(err)
// 		c.JSON(http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, voucher)
// }
