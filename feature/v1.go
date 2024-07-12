package feature

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/transport"
	"net/http"

	"github.com/cesc1802/share-module/system"
	"github.com/gin-gonic/gin"
)

func RegisterHandlerV1(mono system.Service) {
	router := mono.Router()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"data": "success",
		})
	})
	v1 := router.Group("/api/v1")

	auth := v1.Group("/auth")
	{
		auth.POST("/login", transport.Login)

		auth.POST("/register", func(c *gin.Context) {

		})
	}

}
