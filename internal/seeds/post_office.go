package seeds

import (
	"faker/internal/models"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func CreatePostOffice(db *gorm.DB) *models.PostOffice {

	var address models.Address
	if err := db.Order("RANDOM()").First(&address).Error; err != nil {
		log.Println("Failed to fetch a random Address:", err)
		return nil
	}
	log.Printf("Selected Address ID: %d, City: %s, Country: %s", address.ID, address.City, address.Country)

	postOffice := models.PostOffice{
		Name:      gofakeit.Company(),
		Phone:     gofakeit.Phone(),
		AddressID: address.ID,
	}

	if err := db.Create(&postOffice).Error; err != nil {
		log.Println("Failed to insert PostOffice:", err)
		return nil
	}
	log.Printf("Successfully created PostOffice: %+v", postOffice)

	return &postOffice
}
