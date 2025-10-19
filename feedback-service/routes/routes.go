package routes

import (
	"app/controllers"
	"app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	userRoutes := router.Group("/")
	userRoutes.Use(middlewares.AuthMiddleware())
	userRoutes.POST("/feedback", controllers.PostFeedBack)

	adminRoutes := router.Group("/")
	adminRoutes.Use(middlewares.AuthMiddleware())
	adminRoutes.Use(middlewares.AdminMiddleware())
	adminRoutes.GET("/feedback", controllers.GetFeedBack)
	adminRoutes.GET("/feedback/:id", controllers.GetFeedBackById)
	adminRoutes.GET("/feedback/search", controllers.SearchFeedBacks)
	adminRoutes.DELETE("/feedback/:id", controllers.DeleteFeedBack)

}
