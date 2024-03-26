package controllers

import "ecomm/domain/models"

type UserCreateBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserCreateBody) ToModel() *models.User {
	return &models.User{
		Email:    u.Email,
		Password: u.Password,
	}
}

type UserUpdateBody struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *UserUpdateBody) ToModel() *models.User {
	return &models.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}
}
