package routes

import (
	controller "PG-Management-Go-Project/controllers"
	"PG-Management-Go-Project/middleware"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/allpg", controller.Get_All_PG())
	incomingRoutes.GET("/getpgbylocation", controller.Get_PG_By_location())
	incomingRoutes.GET("/getpgbypricemonth", controller.Get_PG_By_Price_Month())
	incomingRoutes.GET("/getpgbypriceday", controller.Get_PG_By_Price_Day())
	incomingRoutes.GET("/getpgbytype", controller.Get_PG_By_Type())
	incomingRoutes.GET("/getpgbyammeneties", controller.Get_PG_By_Ammeneties())
	incomingRoutes.GET("/seebookingbycutomerid", controller.See_booking_by_customerid())
	incomingRoutes.POST("/bookpg", controller.Book_PG())
	incomingRoutes.PATCH("/updatebooking", controller.Update_booking())
	incomingRoutes.DELETE("/canclebooking", controller.Delete_booking())

}
