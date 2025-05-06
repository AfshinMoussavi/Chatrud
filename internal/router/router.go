package router

import "github.com/gin-gonic/gin"

type IRouter interface {
	SetupRoutes(router *gin.RouterGroup)
}
