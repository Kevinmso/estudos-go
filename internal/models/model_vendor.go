package models

type Vendor struct {
	VendorId string `gorm:"primary key" json:"vendorId"`
	Name     string `json:"name"`
	Contact  string `json:"contact"`
	Address  string `json:"address"`
}
