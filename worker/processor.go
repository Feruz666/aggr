package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Feruz666/aggr/models"
	"github.com/Feruz666/aggr/repository"
	"github.com/hibiken/asynq"
)

type TaskProcessor struct {
	Store *repository.Store
}

func NewTaskProcessor(store *repository.Store) *TaskProcessor {
	return &TaskProcessor{
		Store: store,
	}
}

func (m *TaskProcessor) HandleMessageSendTask(ctx context.Context, t *asynq.Task) error {
	pLoad := &models.UsrData{}

	if err := json.Unmarshal(t.Payload(), &pLoad); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	usrData, err := m.Store.Data.CreateData(pLoad)
	if err != nil {
		return fmt.Errorf("CreateData failed: %w", err)
	}

	fmt.Println(usrData)

	return nil
}
