package repositories

import (
	"ecomm/domain/models"
	"ecomm/pkg/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductVariationRepository struct {
	DB *gorm.DB
}

func InitProductVariationRepo() *ProductVariationRepository {
	db, err := database.DeliverDatabaseConnection()
	if err != nil {
		return nil
	}
	return &ProductVariationRepository{
		DB: db,
	}
}

type IProductVariationRepository interface {
	List() ([]models.ProductVariation, error)
	ListByProductID(categoryID int) ([]models.ProductVariation, error)
	FindByID(id int) (*models.ProductVariation, error)
	Create(product *models.ProductVariation) (*models.ProductVariation, error)
	Update(product *models.ProductVariation) (*models.ProductVariation, error)
}

func (r *ProductVariationRepository) List() ([]models.ProductVariation, error) {
	var products []models.ProductVariation
	err := r.DB.Preload(clause.Associations).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductVariationRepository) ListByProductID(productID int) ([]models.ProductVariation, error) {
	var products []models.ProductVariation
	err := r.DB.Preload(clause.Associations).Where("product_id = ?", productID).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductVariationRepository) FindByID(id int) (*models.ProductVariation, error) {
	var product models.ProductVariation
	err := r.DB.Preload(clause.Associations).First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductVariationRepository) Create(product *models.ProductVariation) (*models.ProductVariation, error) {
	err := r.DB.Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductVariationRepository) Update(product *models.ProductVariation) (*models.ProductVariation, error) {
	err := r.DB.Save(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}
