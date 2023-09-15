package repository

import (
	"fmt"
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

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.DbName, cfg.Database.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
