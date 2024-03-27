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
		ParentID: &c.ParentID,
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
		ParentID: &c.ParentID,
	}
}

type ProductVariationCreateBody struct {
	Name      string  `json:"name"`
	Variation string  `json:"variation"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
}

type ProductCreateBody struct {
	Name              string                       `json:"name"`
	Description       string                       `json:"description"`
	CategoryID        uint                         `json:"category_id"`
	Images            string                       `json:"images"`
	Visible           bool                         `json:"visible"`
	ProductVariations []ProductVariationCreateBody `json:"product_variations"`
}

func (p *ProductCreateBody) ToModel() *models.Product {
	var variations []models.ProductVariation
	for _, v := range p.ProductVariations {
		variations = append(variations, models.ProductVariation{
			Name:      v.Name,
			Variation: v.Variation,
			Price:     v.Price,
			Stock:     v.Stock,
		})
	}
	return &models.Product{
		Name:              p.Name,
		Description:       p.Description,
		CategoryID:        p.CategoryID,
		Images:            p.Images,
		Visible:           p.Visible,
		ProductVariations: variations,
	}
}

type ProductVariationUpdateBody struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Variation string  `json:"variation"`
	Price     float64 `json:"price"`
	Stock     int     `json:"stock"`
}

type ProductUpdateBody struct {
	ID                uint                         `json:"id"`
	Name              string                       `json:"name"`
	Description       string                       `json:"description"`
	CategoryID        uint                         `json:"category_id"`
	Images            string                       `json:"images"`
	Visible           bool                         `json:"visible"`
	ProductVariations []ProductVariationUpdateBody `json:"product_variations"`
}

func (p *ProductUpdateBody) ToModel() *models.Product {
	var variations []models.ProductVariation
	for _, v := range p.ProductVariations {
		variations = append(variations, models.ProductVariation{
			ID:        v.ID,
			Name:      v.Name,
			Variation: v.Variation,
			Price:     v.Price,
			Stock:     v.Stock,
		})
	}
	return &models.Product{
		ID:                p.ID,
		Name:              p.Name,
		Description:       p.Description,
		CategoryID:        p.CategoryID,
		Images:            p.Images,
		Visible:           p.Visible,
		ProductVariations: variations,
	}
}
