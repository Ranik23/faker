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
	positions := []string{"Postal Worker", "Delivery Guy"}
	managerPosition := "Director"

	var postOffice models.PostOffice

	if err := db.Order("RANDOM()").First(&postOffice).Error; err != nil {
		log.Printf("Failed to fetch a random post office: %v", err)
		return nil
	}

	log.Printf("Selected PostOffice ID: %d, Name: %s", postOffice.ID, postOffice.Name)

	var existingManager models.Employee
	if err := db.Where("post_office_id = ? AND position = ?", postOffice.ID, managerPosition).First(&existingManager).Error; err == nil {
		log.Printf("Manager (Director) already exists in PostOffice ID %d. Skipping Director creation.", postOffice.ID)
	} else {
		manager := models.Employee{
			Name:         faker.Name(),
			Position:     managerPosition,
			PostOfficeID: postOffice.ID,
			HireDate:     gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now()),
		}

		if err := db.Create(&manager).Error; err != nil {
			log.Printf("Failed to create Manager (Director) for PostOffice ID %d: %v", postOffice.ID, err)
			return nil
		}

		log.Printf("Successfully created Manager (Director): %+v", manager)
	}

	for _, position := range positions {
		employee := models.Employee{
			Name:         faker.Name(),
			Position:     position,
			PostOfficeID: postOffice.ID,
			HireDate:     gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now()),
		}

		if err := db.Create(&employee).Error; err != nil {
			log.Printf("Failed to create Employee (Position: %s) for PostOffice ID %d: %v", position, postOffice.ID, err)
			continue
		}

		log.Printf("Successfully created Employee (Position: %s): %+v", position, employee)
	}

	return nil
}
