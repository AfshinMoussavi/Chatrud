package main

import (
	"Chat-Websocket/config"
	_ "Chat-Websocket/docs"
	"Chat-Websocket/internal/router"
	"Chat-Websocket/internal/user"
	"Chat-Websocket/internal/ws"
	"Chat-Websocket/monitoring"
	"Chat-Websocket/pkg/dbPkg"
	"Chat-Websocket/pkg/loggerPkg"
	"Chat-Websocket/pkg/redisPkg"
	"Chat-Websocket/pkg/utils"
	"Chat-Websocket/pkg/validatorPkg"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func main() {
	r := gin.Default()

	monitoring.InitMetrics()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	cfgPath := utils.GetConfigPath()
	cfg, err := config.InitConfig(cfgPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	logger := loggerPkg.NewLoggerImpl(cfg)
	logger.InitLogger()

	dbConn, err := dbPkg.NewDatabase(cfg)
	if err != nil {
		logger.Fatal(err.Error())
		return
	}
	defer dbConn.Close()

	redisClient, err := redisPkg.InitRedis(cfg)
	if err != nil {
		logger.Fatal("Redis init error: ", err.Error())
	}

	validator := validatorPkg.NewValidator()

	userRepo := user.NewRepository(dbConn.Queries)
	userSvc := user.NewService(userRepo, logger, validator, redisClient)
	userHandler := user.NewHandler(userSvc, logger)

	chatRepo := ws.NewRepository(dbConn.Queries)
	chatSvc := ws.NewService(chatRepo, logger)
	chatHandler := ws.NewHandler(chatSvc, logger)

	userRouter := router.NewRouterImpl(userHandler, chatHandler)
	api := r.Group("/api")

	router.InitRouter(api, userRouter)

	logger.Info("✅ Redis connected successfully")
	logger.Info("✅ Database connected successfully")
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		logger.Fatal(err.Error())
	}

}
