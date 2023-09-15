package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/Feruz666/aggr/logger"
	"github.com/Feruz666/aggr/models"
	"github.com/Feruz666/aggr/worker"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

type Handler struct {
	client *asynq.Client
	logger *logger.Logger
}

func NewHandler(client *asynq.Client, logger *logger.Logger) *Handler {
	return &Handler{
		client: client,
		logger: logger,
	}
}

func (h *Handler) PostMsg(ctx *gin.Context) {
	var usrData models.UsrData
	if err := ctx.BindJSON(&usrData); err != nil {
		log.Printf("could not bind json: %v", err)
		return
	}

	usrData.UserId = ctx.GetHeader(userIdHeader)
	if usrData.UserId == "" {
		log.Println("UserID empty!")
		return
	}

	task, err := worker.NewMessageSendTask(time.Now(), usrData.UserId, usrData.MsgData)
	if err != nil {
		log.Printf("could not create task: %v", err)
		return
	}

	info, err := h.client.Enqueue(task)
	if err != nil {
		log.Printf("could not enqueue task: %v", err)
		return
	}

	h.logger.Debug().Msgf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	ctx.JSON(http.StatusAccepted, gin.H{
		"code": http.StatusAccepted,
	})
}
