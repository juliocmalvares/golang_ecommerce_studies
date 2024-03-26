package repositories

import (
	"ecomm/domain/models"
	"ecomm/pkg/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	DB *gorm.DB
}

func InitProductRepo() *ProductRepository {
	db, err := database.DeliverDatabaseConnection()
	if err != nil {
		return nil
	}
	return &ProductRepository{
		DB: db,
	}
}

type IProductRepository interface {
	List() ([]models.Product, error)
	ListByCategoryID(categoryID int) ([]models.Product, error)
	FindByID(id int) (*models.Product, error)
	Create(product *models.Product) (*models.Product, error)
	Update(product *models.Product) (*models.Product, error)
}

func (r *ProductRepository) List() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Preload(clause.Associations).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) ListByCategoryID(categoryID int) ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Preload(clause.Associations).Where("category_id = ?", categoryID).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) FindByID(id int) (*models.Product, error) {
	var product models.Product
	err := r.DB.Preload(clause.Associations).First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Create(product *models.Product) (*models.Product, error) {
	err := r.DB.Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) Update(product *models.Product) (*models.Product, error) {
	err := r.DB.Save(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}
