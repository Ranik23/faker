package models

import "time"

type PostOffice struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `gorm:"not null"`
	Phone        string
	Employees    []Employee `gorm:"foreignKey:PostOfficeID"`
	PostalItems  []PostalItem `gorm:"foreignKey:PostOfficeID"`
}

type Customer struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `gorm:"not null"`
	ContactInfo  string
	RegDate      time.Time `gorm:"not null"`
	Addresses    []Address `gorm:"foreignKey:ClientID"`
	SentItems    []PostalItem `gorm:"foreignKey:SenderID"`
	ReceivedItems []PostalItem `gorm:"foreignKey:RecipientID"`
	Payments     []Payment `gorm:"foreignKey:CustomerID"`
}

type Address struct {
	ID        uint   `gorm:"primaryKey"`
	Country   string `gorm:"not null"`
	City      string `gorm:"not null"`
	Street    string `gorm:"not null"`
	ClientID  uint   `gorm:"not null"`
	PostalCode string
	Client    Customer `gorm:"foreignKey:ClientID"`
}

type PostalItem struct {
	ID           uint      `gorm:"primaryKey"`
	TrackNum     string    `gorm:"not null;unique"`
	Type         string    `gorm:"not null"`
	Weight       float64   `gorm:"not null"`
	Status       string
	DispDate     time.Time `gorm:"not null"`
	ArrivalDate  time.Time
	SenderID     uint
	RecipientID  uint
	PostOfficeID uint
	Sender       Customer  `gorm:"foreignKey:SenderID"`
	Recipient    Customer  `gorm:"foreignKey:RecipientID"`
	PostOffice   PostOffice `gorm:"foreignKey:PostOfficeID"`
	Payments     []Payment  `gorm:"foreignKey:PostalItemID"`
	StatusTransactions []StatusTransaction `gorm:"foreignKey:PostalItemID"`
}

type Payment struct {
	ID            uint      `gorm:"primaryKey"`
	Amount        float64   `gorm:"not null"`
	PaymentMethod string
	PaymentDate   time.Time `gorm:"not null"`
	PostalItemID  uint
	CustomerID    uint
	PostalItem    PostalItem `gorm:"foreignKey:PostalItemID"`
	Customer      Customer   `gorm:"foreignKey:CustomerID"`
}

type PostalStatus struct {
	ID          uint   `gorm:"primaryKey"`
	StatusName  string `gorm:"not null"`
	Description string
	StatusTransactions []StatusTransaction `gorm:"foreignKey:StatusID"`
}

type Employee struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Position    string    `gorm:"not null"`
	PostOfficeID uint
	HireDate    time.Time `gorm:"not null"`
	PostOffice  PostOffice `gorm:"foreignKey:PostOfficeID"`
}

type StatusTransaction struct {
	ID              uint      `gorm:"primaryKey"`
	PostalItemID    uint      `gorm:"not null"`
	StatusID        uint      `gorm:"not null"`
	StatusDate      time.Time `gorm:"not null"`
	Description     string
	EmployeeID      uint
	PostalItem      PostalItem  `gorm:"foreignKey:PostalItemID"`
	PostalStatus    PostalStatus `gorm:"foreignKey:StatusID"`
	Employee        Employee    `gorm:"foreignKey:EmployeeID"`
}