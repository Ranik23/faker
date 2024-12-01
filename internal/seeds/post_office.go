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

	var companyName string
	isUnique := false
	maxAttempts := 10

	for attempt := 0; attempt < maxAttempts && !isUnique; attempt++ {
		companyName = gofakeit.Company() + " " + gofakeit.Adjective()

		var existingPostOffice models.PostOffice
		if err := db.Where("name = ?", companyName).First(&existingPostOffice).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				isUnique = true
			} else {
				log.Println("Error checking for existing PostOffice:", err)
				return nil
			}
		}

		if isUnique {
			log.Printf("Generated unique company name: %s", companyName)
		}
	}

	if !isUnique {
		log.Println("Failed to generate a unique company name after multiple attempts")
		return nil
	}

	postOffice := models.PostOffice{
		Name:      companyName,
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
