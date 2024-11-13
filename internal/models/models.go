package models

import "time"

type PostOffice struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `gorm:"not null"`
	Phone        string
	Employees    []Employee `gorm:"foreignKey:PostOfficeID;constraint:OnDelete:CASCADE"`
	PostalItems  []PostalItem `gorm:"foreignKey:PostOfficeID;constraint:OnDelete:CASCADE"`
}

type Customer struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `gorm:"not null"`
	ContactInfo  string
	RegDate      time.Time `gorm:"not null"`
	Addresses    []Address `gorm:"foreignKey:ClientID;constraint:OnDelete:CASCADE"`
	SentItems    []PostalItem `gorm:"foreignKey:SenderID;constraint:OnDelete:CASCADE"`
	ReceivedItems []PostalItem `gorm:"foreignKey:RecipientID;constraint:OnDelete:CASCADE"`
	Payments     []Payment `gorm:"foreignKey:CustomerID;constraint:OnDelete:CASCADE"`
}

type Address struct {
	ID        uint   `gorm:"primaryKey"`
	Country   string `gorm:"not null"`
	City      string `gorm:"not null"`
	Street    string `gorm:"not null"`
	ClientID  uint   `gorm:"not null"`
	PostalCode string
	Client    Customer `gorm:"foreignKey:ClientID;constraint:OnDelete:CASCADE"`
}

type PostalItem struct {
	ID           uint      `gorm:"primaryKey"`
	TrackNum     string    `gorm:"not null;unique;index"`
	Type         string    `gorm:"not null"`
	Weight       float64   `gorm:"not null"`
	Status       string
	DispDate     time.Time `gorm:"not null"`
	ArrivalDate  time.Time
	SenderID     uint
	RecipientID  uint
	PostOfficeID uint
	Sender       Customer  `gorm:"foreignKey:SenderID;constraint:OnDelete:CASCADE"`
	Recipient    Customer  `gorm:"foreignKey:RecipientID;constraint:OnDelete:CASCADE"`
	PostOffice   PostOffice `gorm:"foreignKey:PostOfficeID;constraint:OnDelete:CASCADE"`
	Payments     []Payment  `gorm:"foreignKey:PostalItemID;constraint:OnDelete:CASCADE"`
	StatusTransactions []StatusTransaction `gorm:"foreignKey:PostalItemID;constraint:OnDelete:CASCADE"`
}

type Payment struct {
	ID            uint      `gorm:"primaryKey"`
	Amount        float64   `gorm:"not null"`
	PaymentMethod string
	PaymentDate   time.Time `gorm:"not null"`
	PostalItemID  uint
	CustomerID    uint
	PostalItem    PostalItem `gorm:"foreignKey:PostalItemID;constraint:OnDelete:CASCADE"`
	Customer      Customer   `gorm:"foreignKey:CustomerID;constraint:OnDelete:CASCADE"`
}

type PostalStatus struct {
	ID          uint   `gorm:"primaryKey"`
	StatusName  string `gorm:"not null"`
	Description string
	StatusTransactions []StatusTransaction `gorm:"foreignKey:StatusID;constraint:OnDelete:CASCADE"`
}

type Employee struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Position    string    `gorm:"not null"`
	PostOfficeID uint
	HireDate    time.Time `gorm:"not null"`
	PostOffice  PostOffice `gorm:"foreignKey:PostOfficeID;constraint:OnDelete:CASCADE"`
}

type StatusTransaction struct {
	ID              uint      `gorm:"primaryKey"`
	PostalItemID    uint      `gorm:"not null"`
	StatusID        uint      `gorm:"not null"`
	StatusDate      time.Time `gorm:"not null"`
	Description     string
	EmployeeID      uint
	PostalItem      PostalItem  `gorm:"foreignKey:PostalItemID;constraint:OnDelete:CASCADE"`
	PostalStatus    PostalStatus `gorm:"foreignKey:StatusID;constraint:OnDelete:CASCADE"`
	Employee        Employee    `gorm:"foreignKey:EmployeeID;constraint:OnDelete:CASCADE"`
}
