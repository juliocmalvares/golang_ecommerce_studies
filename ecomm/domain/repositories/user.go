package repositories

import (
	"ecomm/domain/models"
	"ecomm/pkg/database"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func InitUserRepo() *UserRepository {
	db, err := database.DeliverDatabaseConnection()
	if err != nil {
		return nil
	}
	return &UserRepository{
		DB: db,
	}
}

type IUserRepository interface {
	FindByID(id int) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
}

func (r *UserRepository) FindByID(id int) (*models.User, error) {
	user := &models.User{}
	err := r.DB.First(user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	var counter int64
	r.DB.Model(&models.User{}).Where("email = ?", user.Email).Count(&counter)
	if counter > 0 {
		return nil, gorm.ErrDuplicatedKey
	}
	err := r.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Update(user *models.User) (*models.User, error) {
	err := r.DB.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
