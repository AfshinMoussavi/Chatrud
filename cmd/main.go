package main

import (
	"Chat-Websocket/config"
	_ "Chat-Websocket/docs"
	"Chat-Websocket/internal/router"
	"Chat-Websocket/internal/user"
	"Chat-Websocket/pkg/dbPkg"
	"Chat-Websocket/pkg/loggerPkg"
	"Chat-Websocket/pkg/redisPkg"
	"Chat-Websocket/pkg/utils"
	"Chat-Websocket/pkg/validatorPkg"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func main() {
	r := gin.Default()

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

	if err := redisPkg.InitRedis(cfg); err != nil {
		logger.Fatal("Redis init error: ", err.Error())
	}

	validator := validatorPkg.NewValidator()

	userRepo := user.NewRepository(dbConn.Queries)
	userSvc := user.NewService(userRepo, logger, validator)
	userHandler := user.NewHandler(userSvc, logger)

	api := r.Group("/api")

	router.InitRouter(api, userHandler)

	logger.Info("✅ Redis connected successfully")
	logger.Info("✅ Database connected successfully")
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		logger.Fatal(err.Error())
	}

}
