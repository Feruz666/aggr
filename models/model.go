package models

import (
	"time"

	"gorm.io/gorm"
)

type UsrData struct {
	gorm.Model
	MsgGetTime time.Time
	UserId     string
	MsgData    Data `json:"data" gorm:"embedded"`
}

type Data struct {
	Headers string `json:"headers"`
	Body    string `json:"body"`
}
