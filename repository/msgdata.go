package repository

import (
	"github.com/Feruz666/aggr/models"
	"gorm.io/gorm"
)

type MsgDataPgRepo struct {
	db *gorm.DB
}

func NewMsgDataPgRepo(db *gorm.DB) *MsgDataPgRepo {
	return &MsgDataPgRepo{
		db: db,
	}
}

func (m *MsgDataPgRepo) CreateData(data *models.UsrData) (*models.UsrData, error) {

	// Create
	res := m.db.Create(&models.UsrData{MsgGetTime: data.MsgGetTime, UserId: data.UserId, MsgData: data.MsgData})
	if res.Error != nil {
		return nil, res.Error
	}

	return data, nil
}
