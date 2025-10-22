package models

type Customer struct {
	CustomerId int    `gorm:"primary_key" json:"customerId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
}
