package router

import (
	"Chat-Websocket/internal/middleware"
	"Chat-Websocket/internal/user"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.RouterGroup, userHandler *user.Handler) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", userHandler.CreateUser)
		auth.GET("/users", userHandler.ListUser)
		auth.POST("/login", userHandler.LoginUser)
	}

	authorized := auth.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.PUT("/edit", userHandler.EditUser)
	}
}
