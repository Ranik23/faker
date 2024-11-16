package seeds

import (
	"faker/internal/models"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func CreatePostalStatus(db *gorm.DB) *models.PostalStatus {
	postalStatus := models.PostalStatus{
		StatusName:  random("Shipped", "Shipping", "Delivered", "Canceled"),
		Description: gofakeit.Sentence(10),
	}

	if err := db.Create(&postalStatus).Error; err != nil {
		log.Println("Failed to insert PostalStatus:", err)
		return nil
	}

	log.Printf("Successfully created PostalStatus: %+v", postalStatus)
	return &postalStatus
}
