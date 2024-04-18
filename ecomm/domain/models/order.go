package models

import "time"

const (
	OrderStatusPending    = "pending"
	OrderStatusProcessing = "processing"
	OrderStatusShipped    = "shipped"
	OrderStatusCompleted  = "completed"
	OrderStatusCancelled  = "cancelled"
)

type Order struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	OrderDate  time.Time   `json:"order_date"`
	TotalPrice float64     `json:"total_price"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
	UserID     uint        `json:"user_id" gorm:"not null"`
	Status     string      `json:"status"` // pending, processing, completed, cancelled
	CreateDate time.Time   `gorm:"autoCreateTime:true; default: now()"`
	UpdateDate time.Time   `gorm:"autoUpdateTime:true"`
}
