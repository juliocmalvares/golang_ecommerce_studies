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

type CategoryCreateBody struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	ParentID uint   `json:"parent_id"`
}

func (c *CategoryCreateBody) ToModel() *models.Category {
	return &models.Category{
		Name:     c.Name,
		Code:     c.Code,
		ParentID: c.ParentID,
	}
}

type CategoryUpdateBody struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	ParentID uint   `json:"parent_id"`
}

func (c *CategoryUpdateBody) ToModel() *models.Category {
	return &models.Category{
		ID:       c.ID,
		Name:     c.Name,
		Code:     c.Code,
		ParentID: c.ParentID,
	}
}
