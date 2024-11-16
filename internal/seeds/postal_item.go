package seeds

import (
	"faker/internal/models"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/bxcodec/faker/v4"
	"gorm.io/gorm"
)
func CreatePostalItem(db *gorm.DB) *models.PostalItem {
	var sender, recipient models.Customer
	var addressSender models.Address
	var postOffice models.PostOffice

	if err := fetchRandomSenderAndPostOffice(db, &sender, &addressSender, &postOffice); err != nil {
		log.Println("Error fetching sender/post office:", err)
		return nil
	}

	if err := db.Order("RANDOM()").First(&recipient).Error; err != nil {
		log.Println("Failed to fetch random recipient:", err)
		return nil
	}

	date := gofakeit.DateRange(time.Now().AddDate(-1, 0, 0), time.Now().AddDate(0, 0, -10))

	status := random("Delivered", "In Delivery")
	var arrivalDate *time.Time
	if status == "Delivered" {
		tempDate := gofakeit.DateRange(date, time.Now())
		arrivalDate = &tempDate
	}

	postalItem := models.PostalItem{
		TrackNum:     faker.UUIDDigit(),
		Type:         random("Package", "Box"),
		Weight:       gofakeit.Float64Range(0.5, 20.0),
		Status:       status,
		DispDate:     date,
		ArrivalDate:  arrivalDate,
		SenderID:     sender.ID,
		RecipientID:  recipient.ID,
		PostOfficeID: postOffice.ID,
	}

	if err := db.Create(&postalItem).Error; err != nil {
		log.Println("Failed to create PostalItem:", err)
		return nil
	}

	log.Printf("Successfully created PostalItem: %+v", postalItem)
	return &postalItem
}

func fetchRandomSenderAndPostOffice(db *gorm.DB, sender *models.Customer, addressSender *models.Address, postOffice *models.PostOffice) error {
	if err := db.Order("RANDOM()").First(sender).Error; err != nil {
		return fmt.Errorf("failed to fetch random sender: %w", err)
	}

	if err := db.Where("customer_id = ?", sender.ID).First(addressSender).Error; err != nil {
		return fmt.Errorf("failed to fetch address for sender: %w", err)
	}

	if err := db.Where("address_id = ?", addressSender.ID).First(postOffice).Error; err != nil {
		return fmt.Errorf("failed to fetch post office for sender's address: %w", err)
	}

	return nil
}


func random(opts ...string) string {
	return opts[rand.Int()%len(opts)]
}
