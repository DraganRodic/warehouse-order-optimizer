package model

type Product struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	SKU      string `json:"sku"`
	Location string `json:"location"`
}