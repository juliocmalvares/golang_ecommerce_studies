package models

import "time"

type Category struct {
	ID         uint       `gorm:"primaryKey"`
	Name       string     `json:"name" gorm:"type:varchar(100);not null"`
	Code       string     `json:"code" gorm:"type:varchar(100);not null"`
	ParentID   uint       `json:"parent_id" gorm:"default:null"`
	Parent     *Category  `json:"parent" gorm:"foreignKey:ParentID;references:ID"`
	Childrens  []Category `json:"childrens" gorm:"foreignKey:ParentID;references:ID"`
	Products   []Product  `json:"products" gorm:"foreignKey:CategoryID;references:ID"`
	CreateDate time.Time  `gorm:"autoCreateTime:true; default: now()"`
	UpdateDate time.Time  `gorm:"autoUpdateTime:true"`
}
