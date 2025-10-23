package models

type Service struct {
	ServiceId string  `gorm:"primary key" json:"service_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
}
