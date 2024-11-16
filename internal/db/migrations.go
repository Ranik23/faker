package db

import (
	"faker/internal/models"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.PostOffice{},
		&models.Customer{},
		&models.Address{},
		&models.PostalItem{},
		&models.Payment{},
		&models.PostalStatus{},
		&models.Employee{},
		&models.StatusTransaction{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
