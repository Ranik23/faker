package main

import (
	"faker/internal/db"
	"faker/internal/seeds"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.Migrate(dbConn)

	for i := 0; i < 1000; i++ {
		seeds.CreateCustomer(dbConn)
	}

	for i := 0; i < 1000; i++ {
		seeds.CreateAddress(dbConn)
	}

	for i := 0; i < 1000; i++ {
		seeds.CreatePostOffice(dbConn)
	}

	for i := 0; i < 1000; i++ {
		seeds.CreatePostalStatus(dbConn)
	}

	for i := 0; i < 1000; i++ {
		seeds.CreateEmployee(dbConn)
	}

	for i := 0; i < 1000; i++ {
		seeds.CreatePostalItem(dbConn)
	}

	for i := 0; i < 1000; i++ {
		seeds.CreatePayment(dbConn)
	}

	for i := 0; i < 1000; i++ {
		seeds.CreateStatusTransaction(dbConn)
	}

	log.Println("Data generation completed successfully!")
}
