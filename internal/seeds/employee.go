package seeds

import (
	"faker/internal/models"
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/bxcodec/faker/v4"
	"gorm.io/gorm"
)

func CreateEmployee(db *gorm.DB) *models.Employee {

	var postOffice models.PostOffice

	for {
		if err := db.Order("RANDOM()").First(&postOffice).Error; err != nil {
			log.Printf("Failed to fetch a random post office: %v", err)
			continue
		}
		break
	}
	log.Printf("Selected PostOffice ID: %d, Name: %s", postOffice.ID, postOffice.Name)

	employee := models.Employee{
		Name:         faker.Name(),
		Position:     random("Postal Worker", "Director", "Delivery Guy"),
		PostOfficeID: postOffice.ID,
		HireDate:     gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now()),
	}

	if err := db.Create(&employee).Error; err != nil {
		log.Printf("Failed to create Employee for PostOffice ID %d: %v", postOffice.ID, err)
		return nil
	}

	log.Printf("Successfully created Employee: %+v", employee)
	return &employee
}

