package routes

import (
	"blue_common/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {
	blueFerry := router.Group("api")
	{
		blueFerry.GET("/health", handler.HealthCheck)
		v1 := blueFerry.Group("v1")
		{
			v1.POST("/register", handler.UserRegister)

		}
	}
}
