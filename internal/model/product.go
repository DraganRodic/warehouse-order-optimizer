package model

type Product struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	SKU      string `json:"sku" gorm:"not null"`
	Location string `json:"location" gorm:"not null"`
}