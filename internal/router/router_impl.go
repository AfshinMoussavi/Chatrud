package router

import (
	"Chat-Websocket/internal/middleware"
	"Chat-Websocket/internal/user"
	"github.com/gin-gonic/gin"
)

type routerImpl struct {
	userHandler user.IUserHandler
}

func NewRouterImpl(userHandler user.IUserHandler) IRouter {
	return &routerImpl{
		userHandler: userHandler,
	}
}

func (r *routerImpl) SetupRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/register", r.userHandler.CreateUserHandler)
		auth.GET("/users", r.userHandler.ListUserHandler)
		auth.POST("/login", r.userHandler.LoginUserHandler)
	}

	authorized := auth.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.PUT("/edit", r.userHandler.EditUserHandler)
	}

	//chat := router.Group("/chat")
	//{
	//	chat.POST("/ws/createRoom", wsHandler.CreateRoom)
	//	chat.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	//	chat.GET("/ws/getRooms", wsHandler.GetRooms)
	//	chat.GET("/ws/getClients/:roomId", wsHandler.GetClients)
	//}
}

func InitRouter(router *gin.RouterGroup, userRouter IRouter) {
	userRouter.SetupRoutes(router)
}
