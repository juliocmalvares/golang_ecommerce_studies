package services

import (
	"ecomm/domain/models"
	"ecomm/domain/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func InitUserService() *UserService {
	repo := repositories.InitUserRepo()
	if repo == nil {
		return nil
	}
	return &UserService{
		userRepository: *repo,
	}
}

type IUserService interface {
	FindByID(id int) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
}

func (s *UserService) FindByID(id int) (*models.User, error) {
	return s.userRepository.FindByID(id)
}

func (s *UserService) Create(user *models.User) (*models.User, error) {
	return s.userRepository.Create(user)
}

func (s *UserService) Update(user *models.User) (*models.User, error) {
	return s.userRepository.Update(user)
}
