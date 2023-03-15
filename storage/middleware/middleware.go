package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/kaolnwza/muniverse/storage/logs"
	"github.com/kaolnwza/muniverse/storage/pkg"
)

func NewGinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := pkg.TokenValid(c)
		if err != nil {
			log.Error(fmt.Errorf("middleware valid token err: %s", err.Error()))
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		claims, err := pkg.ExtractToken(token)
		if err != nil {
			log.Error(fmt.Errorf("middleware extract token err: %s", err.Error()))
			c.JSON(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}

		ctxUserUUID, err := uuid.Parse(fmt.Sprintf("%s", claims["user_uuid"]))
		if err != nil {
			log.Error(fmt.Errorf("middleware parse uuid err: %s", err.Error()))
			c.JSON(http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		ctxKey := "access_user_uuid"
		ctx := context.WithValue(c.Request.Context(), ctxKey, ctxUserUUID)
		c.Request = c.Request.WithContext(ctx)

		c.Next()

	}
}
