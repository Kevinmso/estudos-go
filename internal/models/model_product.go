package models

type Product struct {
	ProductId string  `gorm:"primary key" json:"productId"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	VendorId  string  `json:"vendorId"`
}
