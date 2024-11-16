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

	senderID := postalItem.SenderID
	recipientID := postalItem.RecipientID

	payment := models.Payment{
		Amount:        gofakeit.Float64Range(10.0, 1000.0),
		PaymentMethod: random("Credit Card", "Cash"),
		PaymentDate:   gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now()),
		PostalItemID:  postalItem.ID,
		CustomerID:    senderID,
	}

	payment2 := models.Payment{
		Amount:        gofakeit.Float64Range(10.0, 1000.0),
		PaymentMethod: random("Credit Card", "Cash"),
		PaymentDate:   gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now()),
		PostalItemID:  postalItem.ID,
		CustomerID:    recipientID,
	}

	if err := db.Create(&payment).Error; err != nil {
		log.Printf("Failed to create payment for SenderID %d: %v", senderID, err)
		return nil
	}
	log.Printf("Successfully created payment for SenderID %d: %+v", senderID, payment)

	if err := db.Create(&payment2).Error; err != nil {
		log.Printf("Failed to create payment for RecipientID %d: %v", recipientID, err)
		return nil
	}
	log.Printf("Successfully created payment for RecipientID %d: %+v", recipientID, payment2)

	return &payment
}

