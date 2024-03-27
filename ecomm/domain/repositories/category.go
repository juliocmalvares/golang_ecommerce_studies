package repositories

import (
	"ecomm/domain/models"
	"ecomm/pkg/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func InitCategoryRepo() *CategoryRepository {
	db, err := database.DeliverDatabaseConnection()
	if err != nil {
		return nil
	}
	return &CategoryRepository{
		DB: db,
	}
}

type ICategoryRepository interface {
	List() ([]models.Category, error)
	FindByName(name string) (*models.Category, error)
	FindByID(id uint) (*models.Category, error)
	Create(product *models.Category) (*models.Category, error)
	Update(product *models.Category) (*models.Category, error)
}

func (r *CategoryRepository) Create(category *models.Category) (*models.Category, error) {
	var counter int64
	if category.ParentID == 0 {
		err := r.DB.Model(&models.Category{}).Where("name = ?", category.Name).Count(&counter).Error
		if err != nil {
			return nil, err
		}
		item := models.Category{}
		if counter > 0 {
			r.DB.Find(&item, "name = ?", category.Name)
			return &item, nil
		}
		r.DB.Create(&category)
		return category, nil
	} else {
		var counter int64
		err := r.DB.Model(&models.Category{}).Where("id = ?", category.ParentID).Count(&counter).Error
		if err != nil {
			return nil, err
		}

		parent := models.Category{}
		if counter > 0 {
			r.DB.Find(&parent, "id = ?", category.ParentID)
		} else {
			return nil, err
		}
		category.Parent = &parent
		r.DB.Create(&category)
		return category, nil
	}
}

func (r *CategoryRepository) Update(category *models.Category) (*models.Category, error) {
	var counter int64
	err := r.DB.Model(&models.Category{}).Where("id = ?", category.ID).Count(&counter).Error
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	if counter > 0 {
		r.DB.Save(&category)
		return category, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (r *CategoryRepository) List() ([]models.Category, error) {
	var categories []models.Category
	r.DB.Preload(clause.Associations).Find(&categories)
	return categories, nil
}

func (r *CategoryRepository) FindByName(name string) (*models.Category, error) {
	var category models.Category
	r.DB.Preload(clause.Associations).Find(&category, "name = ?", name)
	if category.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &category, nil
}

func (r *CategoryRepository) FindByID(id uint) (*models.Category, error) {
	var category models.Category
	r.DB.Preload(clause.Associations).Find(&category, "id = ?", id)
	if category.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &category, nil
}
