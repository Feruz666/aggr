package repository

import (
	"fmt"
	"log"

	"github.com/Feruz666/aggr/models"
	"gorm.io/gorm"
)

type Store struct {
	Pg   *gorm.DB
	Data MsgData
}

func NewStore() (*Store, error) {
	// connect to Postgres
	pgDB, err := Dial()
	if err != nil {
		return nil, fmt.Errorf("pg Dial() failed: %w", err)
	}

	if pgDB != nil {
		log.Println("Running Postgres migrations...")
		// Migrate the schemas
		if err := pgDB.AutoMigrate(&models.UsrData{}); err != nil {
			return nil, fmt.Errorf("AutoMigrate UsrData failed: %w", err)
		}
	}

	return &Store{
		Pg:   pgDB,
		Data: NewMsgDataPgRepo(pgDB),
	}, nil
}
