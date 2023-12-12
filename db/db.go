package db

import (
	"log"

	"github.com/atharv-bhadange/log-ingestor/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func Connect() error {

	dsn := "host=localhost user=postgres password=mypass port=5432 sslmode=disable"

	var err error

	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	log.Println("Database connected")

	err = Db.AutoMigrate(&model.Log{})
	if err != nil {
		return err
	}

	log.Println("Database migrated")

	return nil
}
