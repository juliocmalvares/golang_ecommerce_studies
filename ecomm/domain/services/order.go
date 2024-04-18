package services

import (
	"ecomm/domain/models"
	"ecomm/domain/repositories"
)

type OrderService struct {
	OrderRepository     repositories.OrderRepository
	OrderItemRepository repositories.OrderItemRepository
}

func InitOrderService() *OrderService {
	OrderRepo := repositories.InitOrderRepo()
	OrderItemRepo := repositories.InitOrderItemRepo()

	if OrderRepo == nil || OrderItemRepo == nil {
		return nil
	}
	return &OrderService{
		OrderRepository:     *OrderRepo,
		OrderItemRepository: *OrderItemRepo,
	}
}

type IOrderService interface {
	List() ([]models.Order, error)
	ListByUserID(id uint) ([]models.Order, error)
	FindByID(id uint) (*models.Order, error)
	Create(order *models.Order) (*models.Order, error)
	Update(order *models.Order) (*models.Order, error)
}

func (s *OrderService) List() ([]models.Order, error) {
	return s.OrderRepository.List()
}

func (s *OrderService) ListByUserID(id uint) ([]models.Order, error) {
	return s.OrderRepository.ListByUserID(id)
}

func (s *OrderService) FindByID(id uint) (*models.Order, error) {
	return s.OrderRepository.FindByID(id)
}

func (s *OrderService) Create(order *models.Order) (*models.Order, error) {
	ord := models.Order{
		OrderDate: order.OrderDate,
		UserID:    order.UserID,
		Status:    models.OrderStatusPending,
	}
	for _, item := range order.OrderItems {
		ord.TotalPrice += item.Price * float64(item.Quantity)
	}
	insertedOrder, err := s.OrderRepository.Create(&ord)
	if err != nil {
		return nil, err
	}
	for _, item := range order.OrderItems {
		item.OrderID = insertedOrder.ID
		_, err := s.OrderItemRepository.Create(&item)
		if err != nil {
			return nil, err
		}
	}

	return s.OrderRepository.FindByID(insertedOrder.ID)
}

func (s *OrderService) Update(order *models.Order) (*models.Order, error) {
	ord := models.Order{
		ID:         order.ID,
		OrderDate:  order.OrderDate,
		TotalPrice: order.TotalPrice,
		UserID:     order.UserID,
		Status:     order.Status,
	}
	updatedOrder, err := s.OrderRepository.Update(&ord)
	if err != nil {
		return nil, err
	}
	return updatedOrder, nil
}
