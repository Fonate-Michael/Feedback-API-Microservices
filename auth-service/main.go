package main

import (
	"app/db"
	"app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	db.SeedAdmin()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8002")
}
