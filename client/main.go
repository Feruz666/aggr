package main

import (
	"log"

	"github.com/Feruz666/aggr/config"
	"github.com/Feruz666/aggr/handlers"
	"github.com/Feruz666/aggr/logger"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	logger := logger.NewLogger()

	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: cfg.Rds.RedisAddr,
	})

	defer client.Close()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(handlers.UserIdMiddleware())

	h := handlers.NewHandler(client, logger)

	router.POST("/analitycs", h.PostMsg)

	router.Run(cfg.Server.ServerAddr)

}
