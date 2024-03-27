package services

import (
	"ecomm/domain/models"
	"ecomm/domain/repositories"
)

type ProductService struct {
	productRepository          repositories.ProductRepository
	productVariationRepository repositories.ProductVariationRepository
}

func InitProductService() *ProductService {
	productRepo := repositories.InitProductRepo()
	productVariationRepo := repositories.InitProductVariationRepo()
	if productRepo == nil || productVariationRepo == nil {
		return nil
	}
	return &ProductService{
		productRepository:          *productRepo,
		productVariationRepository: *productVariationRepo,
	}
}

type IProductService interface {
	List() ([]models.Product, error)
	FindByName(name string) (*models.Product, error)
	FindByID(id uint) (*models.Product, error)
	ListByCategoryID(categoryID uint) ([]models.Product, error)
	Create(product *models.Product) (*models.Product, error)
	Update(product *models.Product) (*models.Product, error)
}

func (s *ProductService) List() ([]models.Product, error) {
	return s.productRepository.List()
}

func (s *ProductService) FindByName(name string) (*models.Product, error) {
	return s.productRepository.FindByName(name)
}

func (s *ProductService) FindByID(id uint) (*models.Product, error) {
	return s.productRepository.FindByID(id)
}

func (s *ProductService) ListByCategoryID(categoryID uint) ([]models.Product, error) {
	return s.productRepository.ListByCategoryID(categoryID)
}

func (s *ProductService) Create(product *models.Product) (*models.Product, error) {
	prd := models.Product{
		Name:        product.Name,
		Description: product.Description,
		Visible:     product.Visible,
		Images:      product.Images,
		CategoryID:  product.CategoryID,
	}
	insertedProduct, err := s.productRepository.Create(&prd)
	if err != nil {
		return nil, err
	}
	variations := make([]models.ProductVariation, 0)
	for _, v := range product.ProductVariations {
		variations = append(variations, models.ProductVariation{
			ProductID: insertedProduct.ID,
			Name:      v.Name,
			Variation: v.Variation,
			Price:     v.Price,
			Stock:     v.Stock,
		})
	}
	for _, v := range variations {
		_, err := s.productVariationRepository.Create(&v)
		if err != nil {
			return nil, err
		}
	}

	return s.productRepository.FindByID(insertedProduct.ID)
}

func (s *ProductService) Update(product *models.Product) (*models.Product, error) {
	prd := models.Product{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Visible:     product.Visible,
		Images:      product.Images,
		CategoryID:  product.CategoryID,
	}
	updatedProduct, err := s.productRepository.Update(&prd)
	if err != nil {
		return nil, err
	}
	variations := make([]models.ProductVariation, 0)
	for _, v := range product.ProductVariations {
		variations = append(variations, models.ProductVariation{
			ID:        v.ID,
			ProductID: updatedProduct.ID,
			Name:      v.Name,
			Variation: v.Variation,
			Price:     v.Price,
			Stock:     v.Stock,
		})
	}
	for _, v := range variations {
		_, err := s.productVariationRepository.Update(&v)
		if err != nil {
			return nil, err
		}
	}

	return s.productRepository.FindByID(updatedProduct.ID)
}
