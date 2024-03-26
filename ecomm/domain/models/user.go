package models

import "time"

type User struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Email      string    `json:"email" gorm:"unique;required"`
	Password   string    `json:"password" gorm:"required"`
	CreateDate time.Time `gorm:"autoCreateTime:true; default: now()"`
	UpdateDate time.Time `gorm:"autoUpdateTime:true"`
}
