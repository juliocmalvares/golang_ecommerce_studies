package repositories

import (
	"ecomm/domain/models"
	"ecomm/pkg/database"

	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func InitOrderRepo() *OrderRepository {
	db, err := database.DeliverDatabaseConnection()
	if err != nil {
		return nil
	}
	return &OrderRepository{
		DB: db,
	}
}

type IOrderRepository interface {
	List() ([]models.Order, error)
	ListByUserID(id uint) ([]models.Order, error)
	FindByID(id uint) (*models.Order, error)
	Create(order *models.Order) (*models.Order, error)
	Update(order *models.Order) (*models.Order, error)
}

func (r *OrderRepository) List() ([]models.Order, error) {
	var orders []models.Order
	if err := r.DB.Preload("OrderItems.Product").Preload("OrderItems.ProductVariation").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) ListByUserID(id uint) ([]models.Order, error) {
	var orders []models.Order
	if err := r.DB.Preload("OrderItems.Product").Preload("OrderItems.ProductVariation").Where("user_id = ?", id).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) FindByID(id uint) (*models.Order, error) {
	var order models.Order
	if err := r.DB.Preload("OrderItems.Product").Preload("OrderItems.ProductVariation").Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) Create(order *models.Order) (*models.Order, error) {
	if err := r.DB.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepository) Update(order *models.Order) (*models.Order, error) {
	if err := r.DB.Save(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}
