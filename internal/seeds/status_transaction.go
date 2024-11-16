package seeds

import (
	"faker/internal/models"
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func CreateStatusTransaction(db *gorm.DB) *models.StatusTransaction {
	var postalItem models.PostalItem
	var postalStatus models.PostalStatus

	for {
		if err := db.Order("RANDOM()").First(&postalItem).Error; err != nil {
			log.Println("Failed to get random PostalItem:", err)
			continue
		}
		break
	}
	log.Printf("Selected PostalItem: %+v", postalItem)

	for {
		if err := db.Order("RANDOM()").First(&postalStatus).Error; err != nil {
			log.Println("Failed to get random PostalStatus:", err)
			continue
		}
		break
	}
	log.Printf("Selected PostalStatus: %+v", postalStatus)

	var postOffice models.PostOffice
	if err := db.Where("id = ?", postalItem.PostOfficeID).First(&postOffice).Error; err != nil {
		log.Println("Failed to get PostOffice:", err)
		return nil
	}
	log.Printf("Associated PostOffice: %+v", postOffice)

	var employee models.Employee
	if err := db.Where("post_office_id = ?", postOffice.ID).First(&employee).Error; err != nil {
		log.Println("Failed to get Employee:", err)
		return nil
	}
	log.Printf("Assigned Employee: %+v", employee)

	statusTransaction := models.StatusTransaction{
		PostalItemID: postalItem.ID,
		StatusID:     postalStatus.ID,
		StatusDate:   gofakeit.DateRange(time.Now().AddDate(0, -1, 0), time.Now()),
		Description:  gofakeit.Sentence(10),
		EmployeeID:   employee.ID,
	}

	if err := db.Create(&statusTransaction).Error; err != nil {
		log.Println("Failed to insert StatusTransaction:", err)
		return nil
	}

	log.Printf("Successfully created StatusTransaction: %+v", statusTransaction)
	return &statusTransaction
}
