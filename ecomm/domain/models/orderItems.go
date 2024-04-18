package models

import "time"

type OrderItem struct {
	ID                 uint             `json:"id" gorm:"primaryKey"`
	OrderID            uint             `json:"order_id" gorm:"not null"`
	ProductID          uint             `json:"product_id" gorm:"not null"`
	Product            Product          `json:"product" gorm:"foreignKey:ProductID"`
	ProductVariationID uint             `json:"product_variation_id" gorm:"not null"`
	ProductVariation   ProductVariation `json:"product_variation" gorm:"foreignKey:ProductVariationID"`
	Quantity           uint             `json:"quantity" gorm:"not null"`
	Price              float64          `json:"price" gorm:"not null"`
	Subtotal           float64          `json:"subtotal" gorm:"not null"`
	Discount           float64          `json:"discount"`
	CreateDate         time.Time        `gorm:"autoCreateTime:true; default: now()"`
	UpdateDate         time.Time        `gorm:"autoUpdateTime:true"`
}
