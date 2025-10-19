package main

import (
	"app/db"
	"app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8003")
}
