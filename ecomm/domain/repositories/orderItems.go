package repositories

import (
	"ecomm/domain/models"
	"ecomm/pkg/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OrderItemRepository struct {
	DB *gorm.DB
}

func InitOrderItemRepo() *OrderItemRepository {
	db, err := database.DeliverDatabaseConnection()
	if err != nil {
		return nil
	}
	return &OrderItemRepository{
		DB: db,
	}
}

type IOrderItemRepository interface {
	List() ([]models.OrderItem, error)
	FindByID(id uint) (*models.OrderItem, error)
	ListByOrderID(id uint) ([]models.OrderItem, error)
	Create(order *models.OrderItem) (*models.OrderItem, error)
	Update(order *models.OrderItem) (*models.OrderItem, error)
	Remove(order *models.OrderItem) error
}

func (r *OrderItemRepository) List() ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	if err := r.DB.Preload(clause.Associations).Find(&orderItems).Error; err != nil {
		return nil, err
	}
	return orderItems, nil
}

func (r *OrderItemRepository) FindByID(id uint) (*models.OrderItem, error) {
	var orderItem models.OrderItem
	if err := r.DB.Preload(clause.Associations).Where("id = ?", id).First(&orderItem).Error; err != nil {
		return nil, err
	}
	return &orderItem, nil
}

func (r *OrderItemRepository) ListByOrderID(id uint) ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	if err := r.DB.Preload(clause.Associations).Where("order_id = ?", id).Find(&orderItems).Error; err != nil {
		return nil, err
	}
	return orderItems, nil
}

func (r *OrderItemRepository) Create(order *models.OrderItem) (*models.OrderItem, error) {
	if err := r.DB.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderItemRepository) Update(order *models.OrderItem) (*models.OrderItem, error) {
	if err := r.DB.Save(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderItemRepository) Remove(order *models.OrderItem) error {
	if err := r.DB.Delete(order).Error; err != nil {
		return err
	}
	return nil
}
