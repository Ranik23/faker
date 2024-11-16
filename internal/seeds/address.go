package seeds

import (
	"faker/internal/models"
	random_user "faker/internal/random"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func CreateAddress(db *gorm.DB) *models.Address {

	var customer models.Customer
	if err := db.Order("RANDOM()").First(&customer).Error; err != nil {
		log.Printf("Failed to fetch a random customer: %v", err)
		return nil
	}
	log.Printf("Selected customer ID: %d, Name: %s", customer.ID, customer.Name)

	country, city, street := random_user.GetRandomUser()

	if country == "" || city == "" {
		country = gofakeit.Country()
		city = gofakeit.City()
		street = gofakeit.StreetName()
	}

	address := models.Address{
		Country:    country,
		City:       city,
		Street:     street,
		PostalCode: gofakeit.Zip(),
		CustomerID: customer.ID,
	}

	if err := db.Create(&address).Error; err != nil {
		log.Printf("Failed to create address for customer ID %d: %v", customer.ID, err)
		return nil
	}
	log.Printf("Successfully created address: %+v", address)

	return &address
}
