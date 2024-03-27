package models

import "time"

type Product struct {
	ID                uint               `json:"id" gorm:"primaryKey"`
	Name              string             `json:"name"`
	Description       string             `json:"description"`
	Visible           bool               `json:"visible"`
	Images            string             `json:"images"`
	ProductVariations []ProductVariation `json:"product_variation" gorm:"foreignKey:ProductID"`
	CategoryID        uint               `json:"category_id" gorm:"not null"`
	CreateDate        time.Time          `gorm:"autoCreateTime:true; default: now()"`
	UpdateDate        time.Time          `gorm:"autoUpdateTime:true"`
}
