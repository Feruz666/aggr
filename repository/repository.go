package repository

import (
	"github.com/Feruz666/aggr/models"
)

type MsgData interface {
	CreateData(*models.UsrData) (*models.UsrData, error)
}
