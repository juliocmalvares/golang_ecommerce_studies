package models

import "time"

type User struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Email      string    `json:"email" gorm:"unique"`
	Password   string    `json:"password"`
	CreateDate time.Time `gorm:"autoCreateTime:true; default: now()"`
	UpdateDate time.Time `gorm:"autoUpdateTime:true"`
}
