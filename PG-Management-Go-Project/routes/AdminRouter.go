package routes

import (
	controller "PG-Management-Go-Project/controllers"
	"PG-Management-Go-Project/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())

}
