package services

import (
	"ecomm/domain/models"
	"ecomm/domain/repositories"
)

type OrderItemService struct {
	OrderItemRepository repositories.OrderItemRepository
	OrderService        OrderService
}

func InitOrderItemService() *OrderItemService {
	OrderItemRepo := repositories.InitOrderItemRepo()
	OrderService := InitOrderService()

	if OrderItemRepo == nil || OrderService == nil {
		return nil
	}
	return &OrderItemService{
		OrderItemRepository: *OrderItemRepo,
		OrderService:        *OrderService,
	}
}

type IOrderItemService interface {
	List() ([]models.OrderItem, error)
	FindByID(id uint) (*models.OrderItem, error)
	ListByOrderID(id uint) ([]models.OrderItem, error)
	Create(order *models.OrderItem) (*models.OrderItem, error)
	Update(order *models.OrderItem) (*models.OrderItem, error)
	AddOrderItem(orderItem *models.OrderItem) (*models.OrderItem, error)
	RemoveOrderItem(orderItem *models.OrderItem) error
}

func (s *OrderItemService) List() ([]models.OrderItem, error) {
	return s.OrderItemRepository.List()
}

func (s *OrderItemService) FindByID(id uint) (*models.OrderItem, error) {
	return s.OrderItemRepository.FindByID(id)
}

func (s *OrderItemService) ListByOrderID(id uint) ([]models.OrderItem, error) {
	return s.OrderItemRepository.ListByOrderID(id)
}

func (s *OrderItemService) Create(order *models.OrderItem) (*models.OrderItem, error) {
	return s.OrderItemRepository.Create(order)
}

func (s *OrderItemService) Update(order *models.OrderItem) (*models.OrderItem, error) {
	return s.OrderItemRepository.Update(order)
}

func (s *OrderItemService) AddOrderItem(orderItem *models.OrderItem) (*models.OrderItem, error) {
	order, err := s.OrderService.FindByID(orderItem.OrderID)
	if err != nil {
		return nil, err
	}
	order.TotalPrice += orderItem.Price * float64(orderItem.Quantity)
	_, err = s.OrderService.Update(order)
	if err != nil {
		return nil, err
	}
	return s.Create(orderItem)
}

func (s *OrderItemService) RemoveOrderItem(orderItem *models.OrderItem) error {
	order, err := s.OrderService.FindByID(orderItem.OrderID)
	if err != nil {
		return err
	}
	order.TotalPrice -= orderItem.Price * float64(orderItem.Quantity)
	_, err = s.OrderService.Update(order)
	if err != nil {
		return err
	}
	return s.OrderItemRepository.Remove(orderItem)
}
