package routes

import (
	"auth-service/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	userHandler *handler.UserHandler,
) {
	api := router.Group("/v1")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "pong",
		})
	})

	// User (Auth)
	userGroup := api.Group("/user")
	{
		userGroup.POST("/register", userHandler.Register)
		userGroup.POST("/login", userHandler.Login)
	}
}
