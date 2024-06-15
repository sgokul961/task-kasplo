package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/pkg/config"
	"main.go/pkg/models"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)

	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{SkipDefaultTransaction: true})

	// if dbErr != nil {
	// 	return nil, dbErr
	// }

	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.ToDo{})

	return db, dbErr
}
