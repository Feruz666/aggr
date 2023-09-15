package worker

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Feruz666/aggr/models"
	"github.com/hibiken/asynq"
)

const TypeTaskMsg = "task:send_msg"

// NewMessageSendTask  creates a task.
// A task consists of a type and a payload.
func NewMessageSendTask(msgGetTime time.Time, userId string, msgData models.Data) (*asynq.Task, error) {
	payload, err := json.Marshal(
		models.UsrData{
			MsgGetTime: msgGetTime,
			UserId:     userId,
			MsgData:    msgData,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("marshal error %w", err)
	}

	return asynq.NewTask(TypeTaskMsg, payload), nil
}
