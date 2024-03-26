package models

import "time"

type ProductVariation struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	ProductID  uint      `json:"product_id"`
	Product    Product   `json:"product" gorm:"foreignKey:ProductID"`
	Variation  string    `json:"variation"`
	Price      float64   `json:"price"`
	Stock      int       `json:"stock"`
	CreateDate time.Time `gorm:"autoCreateTime:true; default: now()"`
	UpdateDate time.Time `gorm:"autoUpdateTime:true"`
}
