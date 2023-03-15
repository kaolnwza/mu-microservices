package main

import (
	"os"

	"github.com/gin-gonic/gin"
	service "github.com/kaolnwza/muniverse/auth/services"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/oauth", service.GoogleCallackByTokenHandler)
			auth.POST("/test", service.GoogleLoginTestHandler)
		}
	}

	r.Run(":" + os.Getenv("REST_PORT"))
}
