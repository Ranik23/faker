package seeds

import (
	"faker/internal/models"
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/bxcodec/faker/v4"
	"gorm.io/gorm"
)

func CreateCustomer(db *gorm.DB) *models.Customer {
	customer := models.Customer{
		Name:        faker.Name(),
		ContactInfo: gofakeit.Email(),
		RegDate:     gofakeit.DateRange(time.Now().AddDate(-5, 0, 0), time.Now()),
	}

	if err := db.Create(&customer).Error; err != nil {
		log.Printf("Failed to create customer: %v", err)
		return nil
	}

	log.Printf("Successfully created customer: %+v", customer)
	return &customer
}
