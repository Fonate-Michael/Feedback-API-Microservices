package main

import (
	"app/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env: ", err)
		return
	}
	AUTH_SERVICE := os.Getenv("AUTH_SERVICE")
	FEEDBACK_SERIVCE := os.Getenv("FEEDBACK_SERVICE")

	router := gin.Default()

	router.Use(middleware.RateLimiter())
	router.Any("/auth-service/*path", middleware.ReverseProxy(AUTH_SERVICE, "/auth-service"))
	router.Any("/feedback-service/*path", middleware.ReverseProxy(FEEDBACK_SERIVCE, "/feedback-service"))
	router.Run(":8000")
}
