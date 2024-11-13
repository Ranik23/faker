package main

import (
	"faker/internal/models"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/bxcodec/faker/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func random(opts ...string) string {
	if len(opts) == 0 {
		return ""
	}
	return opts[rand.Int()%len(opts)]
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)


	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.PostOffice{}, &models.Customer{}, &models.Address{}, &models.PostalItem{}, &models.Payment{}, &models.PostalStatus{}, &models.Employee{}, &models.StatusTransaction{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	for i := 0; i < 100; i++ {
		postOffice := models.PostOffice{
			Name:  gofakeit.Company(),
			Phone: gofakeit.Phone(),
		}
		if err := db.Create(&postOffice).Error; err != nil {
			log.Println("Failed to insert PostOffice:", err)
			continue
		}

		sender := models.Customer{
			Name:        faker.Name(),
			ContactInfo: gofakeit.Email(),
			RegDate:     gofakeit.Date(),
		}
		recipient := models.Customer{
			Name:        faker.Name(),
			ContactInfo: gofakeit.Email(),
			RegDate:     gofakeit.Date(),
		}
		db.Create(&sender)
		db.Create(&recipient)

		address := models.Address{
			Country:    gofakeit.Country(),
			City:       gofakeit.City(),
			Street:     gofakeit.StreetName(),
			ClientID:   sender.ID,
			PostalCode: gofakeit.Zip(),
		}
		if err := db.Create(&address).Error; err != nil {
			log.Println("Failed to insert Address:", err)
			continue
		}

		postalItem := models.PostalItem{
			TrackNum:     faker.UUIDDigit(),
			Type:         random("Package", "Box"),
			Weight:       gofakeit.Float64Range(0.5, 20.0),
			Status:       random("Sent", "In Progress"),
			DispDate:     time.Now(),
			ArrivalDate:  time.Now().Add(48 * time.Hour),
			SenderID:     sender.ID,
			RecipientID:  recipient.ID,
			PostOfficeID: postOffice.ID,
		}
		if err := db.Create(&postalItem).Error; err != nil {
			log.Println("Failed to insert PostalItem:", err)
			continue
		}

		payment := models.Payment{
			Amount:        gofakeit.Float64Range(10.0, 1000.0),
			PaymentMethod: random("Credit Card", "Cash"),
			PaymentDate:   time.Now(),
			PostalItemID:  postalItem.ID,
			CustomerID:    sender.ID,
		}
		if err := db.Create(&payment).Error; err != nil {
			log.Println("Failed to insert Payment:", err)
			continue
		}

		postalStatus := models.PostalStatus{
			StatusName:  random("Shipped", "Shipping"),
			Description: "Package has been shipped and is on the way.",
		}
		if err := db.Create(&postalStatus).Error; err != nil {
			log.Println("Failed to insert PostalStatus:", err)
			continue
		}

		employee := models.Employee{
			Name:        faker.Name(),
			Position:    random("Postal Worker", "Director", "Delivery Guy"),
			PostOfficeID: postOffice.ID,
			HireDate:    time.Now(),
		}
		if err := db.Create(&employee).Error; err != nil {
			log.Println("Failed to insert Employee:", err)
			continue
		}

		statusTransaction := models.StatusTransaction{
			PostalItemID: postalItem.ID,
			StatusID:     postalStatus.ID,
			StatusDate:   time.Now(),
			Description:  "Item dispatched from warehouse",
			EmployeeID:   employee.ID,
		}
		if err := db.Create(&statusTransaction).Error; err != nil {
			log.Println("Failed to insert StatusTransaction:", err)
			continue
		}
	}
	log.Println("Data generation completed successfully!")
}
