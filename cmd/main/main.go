package main

import (
	"faker/internal/db"
	"faker/internal/seeds"
	"log"
)

func main() {

	dbConn := db.Connect()

	db.Migrate(dbConn)

	for i := 0; i < 10000; i++ {
		seeds.CreateCustomer(dbConn)
	}

	for i := 0; i < 400; i++ {
		seeds.CreateAddress(dbConn)
	}

	for i := 0; i < 10000; i++ {
		seeds.CreatePostOffice(dbConn)
	}

	seeds.CreatePostalStatus(dbConn)

	for i := 0; i < 10000; i++ {
		seeds.CreateEmployee(dbConn)
	}

	for i := 0; i < 10000; i++ {
		seeds.CreatePostalItem(dbConn)
	}

	for i := 0; i < 10000; i++ {
		seeds.CreatePayment(dbConn)
	}

	for i := 0; i < 10000; i++ {
		seeds.CreateStatusTransaction(dbConn)
	}

	log.Println("Data generation completed successfully!")
}
