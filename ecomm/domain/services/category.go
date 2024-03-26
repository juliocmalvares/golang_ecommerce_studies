package services

import (
	"ecomm/domain/models"
	"ecomm/domain/repositories"
)

type CategoryService struct {
	categoryRepository repositories.CategoryRepository
}

func InitCategoryService() *CategoryService {
	repo := repositories.InitCategoryRepo()
	if repo == nil {
		return nil
	}
	return &CategoryService{
		categoryRepository: *repo,
	}
}

type ICategoryService interface {
	List() ([]models.Category, error)
	FindByName(name string) (*models.Category, error)
	FindByID(id uint) (*models.Category, error)
	Create(product *models.Category) (*models.Category, error)
	Update(product *models.Category) (*models.Category, error)
}

func (s *CategoryService) List() ([]models.Category, error) {
	return s.categoryRepository.List()
}

func (s *CategoryService) FindByName(name string) (*models.Category, error) {
	return s.categoryRepository.FindByName(name)
}

func (s *CategoryService) FindByID(id uint) (*models.Category, error) {
	return s.categoryRepository.FindByID(id)
}

func (s *CategoryService) Create(category *models.Category) (*models.Category, error) {
	return s.categoryRepository.Create(category)
}

func (s *CategoryService) Update(category *models.Category) (*models.Category, error) {
	return s.categoryRepository.Update(category)
}
