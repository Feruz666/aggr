package repository

import (
	"log"

	"github.com/Feruz666/aggr/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Dial() (*gorm.DB, error) {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open(postgres.Open(cfg.Database.Url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	
	return db, nil
}
