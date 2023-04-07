package routes

import (
	controller "PG-Management-Go-Project/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/signup", controller.Signup())
	incomingRoutes.POST("/signup", controller.Signup())
	incomingRoutes.GET("/login", controller.Login())
	incomingRoutes.POST("/login", controller.Login())

}
