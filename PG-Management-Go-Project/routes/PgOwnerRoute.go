package routes

import (
	controller "PG-Management-Go-Project/controllers"
	"PG-Management-Go-Project/middleware"

	"github.com/gin-gonic/gin"
)

func PgOwnerRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/addproperty", controller.Add_property())
	incomingRoutes.POST("/addproperty", controller.Add_property())
	incomingRoutes.PATCH("/updateproperty", controller.Update_property())
	incomingRoutes.DELETE("/deleteproperty", controller.Delete_property())
	incomingRoutes.GET("/seebooking", controller.See_bookings())
	incomingRoutes.GET("/seebookingsbypropertyid", controller.See_booking_by_propertyid())

}
