package seeds

import (
	"faker/internal/models"
	"log"
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func CreateStatusTransaction(db *gorm.DB) *models.StatusTransaction {
	var postalItem models.PostalItem

	if err := db.Order("RANDOM()").First(&postalItem).Error; err != nil {
		log.Println("Failed to get random PostalItem:", err)
		return nil
	}
	log.Printf("Selected PostalItem: %+v", postalItem)

	var transactionCount int64
	if err := db.Model(&models.StatusTransaction{}).
		Where("postal_item_id = ?", postalItem.ID).
		Count(&transactionCount).Error; err != nil {
		log.Println("Failed to count StatusTransactions:", err)
		return nil
	}

	if transactionCount >= 4 {
		log.Println("Maximum number of transactions reached for this PostalItem.")
		return nil
	}

	var usedStatusIDs []uint
	if err := db.Model(&models.StatusTransaction{}).
		Where("postal_item_id = ?", postalItem.ID).
		Pluck("status_id", &usedStatusIDs).Error; err != nil {
		log.Println("Failed to get used statuses:", err)
		return nil
	}

	var lastStatus models.StatusTransaction
	if err := db.Where("postal_item_id = ?", postalItem.ID).
		Order("status_date DESC").
		First(&lastStatus).Error; err != nil {
		log.Println("No previous statuses found. Starting with 'Shipped'.")
	} else {
		log.Printf("Last StatusTransaction: %+v", lastStatus)
	}

	var newStatus models.PostalStatus
	var newStatusID uint
	if lastStatus.StatusID == 0 {
		newStatusID = 1 // Shipped
	} else if lastStatus.StatusID == 1 {
		newStatusID = 2 // Shipping
	} else if lastStatus.StatusID == 2 {
		if rand.Float64() < 0.3 {
			newStatusID = 4 // Cancelled
		} else {
			newStatusID = 3 // Delivered
		}
	} else {
		return nil
	}

	if err := db.Where("id = ?", newStatusID).First(&newStatus).Error; err != nil {
		log.Println("Failed to get PostalStatus:", err)
		return nil
	}
	log.Printf("Selected PostalStatus: %+v", newStatus)

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

	var description string

	shippedDescriptions := []string{
		"The item has been shipped and is on its way to you.",
		"Your item is now on the way, it has been shipped.",
		"Shipment has been dispatched. It is now on the way.",
	}
	
	shippingDescriptions := []string{
		"The item is currently in transit and moving towards its destination.",
		"The shipment is in progress. The item is on its way.",
		"Shipping is ongoing, and the item is expected to arrive soon.",
	}
	
	deliveredDescriptions := []string{
		"The item has been successfully delivered to the recipient.",
		"Your item has been delivered and signed for.",
		"Delivery completed. The item has reached its destination.",
	}
	
	cancelledDescriptions := []string{
		"The shipment has been cancelled and will not proceed.",
		"The order has been cancelled and no further actions will be taken.",
		"The item has been cancelled and is no longer in transit.",
	}
	
	switch newStatusID {
	case 1: // Shipped
		description = shippedDescriptions[rand.Intn(len(shippedDescriptions))]
	case 2: // Shipping
		description = shippingDescriptions[rand.Intn(len(shippingDescriptions))]
	case 3: // Delivered
		description = deliveredDescriptions[rand.Intn(len(deliveredDescriptions))]
	case 4: // Cancelled
		description = cancelledDescriptions[rand.Intn(len(cancelledDescriptions))]
	default:
		description = "Status update for the item."
	}

	statusTransaction := models.StatusTransaction{
		PostalItemID: postalItem.ID,
		StatusID:     newStatusID,
		StatusDate:   gofakeit.DateRange(time.Now().AddDate(0, -1, 0), time.Now().Add(24 * time.Hour)),
		Description:  description,
		EmployeeID:   employee.ID,
	}

	if err := db.Create(&statusTransaction).Error; err != nil {
		log.Println("Failed to insert StatusTransaction:", err)
		return nil
	}

	log.Printf("Successfully created StatusTransaction: %+v", statusTransaction)
	return &statusTransaction
}
