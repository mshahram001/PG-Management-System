package main

import (
	database "PG-Management-Go-Project/database"

	routes "PG-Management-Go-Project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	port := "8000"
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("template/*.html")
	// routes.HomeRoutes(router)
	routes.AuthRoutes(router)
	routes.AdminRoutes(router)
	routes.CustomerRoutes(router)
	routes.PgOwnerRoutes(router)

	router.Run(":" + port)

}
