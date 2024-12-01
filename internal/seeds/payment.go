package seeds

import (
	"faker/internal/models"
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func CreatePayment(db *gorm.DB) *models.Payment {

	var postalItem models.PostalItem
	
	if err := db.Order("RANDOM()").First(&postalItem).Error; err != nil {
		log.Println("Failed to fetch a random PostalItem:", err)
		return nil	
	}

	log.Printf("Selected PostalItem ID: %d, TrackNum: %s", postalItem.ID, postalItem.TrackNum)

	var existingPayments []models.Payment
	if err := db.Where("postal_item_id = ?", postalItem.ID).Find(&existingPayments).Error; err != nil {
		log.Println("Failed to check existing payments:", err)
		return nil
	}

	if len(existingPayments) > 0 {
		log.Printf("PostalItem ID %d already has a payment, skipping creation.", postalItem.ID)
		return nil
	}

	senderID := postalItem.SenderID

	payment := models.Payment{
		Amount:        gofakeit.Float64Range(10.0, 1000.0),
		PaymentMethod: random("Credit Card", "Cash"),
		PaymentDate:   gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now()),
		PostalItemID:  postalItem.ID,
		CustomerID:    senderID,
	}

	if err := db.Create(&payment).Error; err != nil {
		log.Printf("Failed to create payment for SenderID %d: %v", senderID, err)
		return nil
	}
	log.Printf("Successfully created payment for SenderID %d: %+v", senderID, payment)

	return &payment
}
