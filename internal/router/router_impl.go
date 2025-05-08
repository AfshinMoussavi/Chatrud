package router

import (
	"Chat-Websocket/internal/middleware"
	"Chat-Websocket/internal/user"
	"Chat-Websocket/internal/ws"
	"github.com/gin-gonic/gin"
)

type routerImpl struct {
	userHandler user.IUserHandler
	chatHandler ws.IWsHandler
}

func NewRouterImpl(userHandler user.IUserHandler, wsHandler ws.IWsHandler) IRouter {
	return &routerImpl{
		userHandler: userHandler,
		chatHandler: wsHandler,
	}
}

func (r *routerImpl) SetupRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", r.userHandler.CreateUserHandler)
		auth.GET("/users", r.userHandler.ListUserHandler)
		auth.POST("/login", r.userHandler.LoginUserHandler)
	}
	auth.GET("/ws", r.chatHandler.HandleWebSocket)

	authorized := auth.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{

		authorized.PUT("/edit", r.userHandler.EditUserHandler)
		authorized.DELETE("/delete", r.userHandler.DeleteUserHandler)
	}
}

func InitRouter(router *gin.RouterGroup, userRouter IRouter) {
	userRouter.SetupRoutes(router)
}
