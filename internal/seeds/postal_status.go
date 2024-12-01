package seeds

import (
	"faker/internal/models"
	"log"
	"gorm.io/gorm"
)

func CreatePostalStatus(db *gorm.DB) *models.PostalStatus {
	postalStatus := models.PostalStatus{
		StatusName:  "Shipped", // "Shipping", "Delivered", "Canceled"),
		Description: "Item Was Shipped",
	}

	postalStatus2 := models.PostalStatus{
		StatusName:  "Shipping", //"Delivered", "Canceled"),
		Description: "Item Is On Your Way",
	}

	postalStatus3 := models.PostalStatus{
		StatusName:  "Delivered",// "Canceled"),
		Description: "Item Is Delivered",
	}

	postalStatus4 := models.PostalStatus{
		StatusName:  "Canceled",
		Description: "Item Is Cancelled",
	}

	if err := db.Create(&postalStatus).Error; err != nil {
		log.Println("Failed to insert PostalStatus:", err)
		return nil
	}

	if err := db.Create(&postalStatus2).Error; err != nil {
		log.Println("Failed to insert PostalStatus:", err)
		return nil
	}

	if err := db.Create(&postalStatus3).Error; err != nil {
		log.Println("Failed to insert PostalStatus:", err)
		return nil
	}

	if err := db.Create(&postalStatus4).Error; err != nil {
		log.Println("Failed to insert PostalStatus:", err)
		return nil
	}


	log.Printf("Successfully created PostalStatus: %+v", postalStatus)
	return &postalStatus
}
