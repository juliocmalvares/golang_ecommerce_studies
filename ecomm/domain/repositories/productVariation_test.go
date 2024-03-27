package repositories

import (
	"ecomm/domain/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductVariationRepository_Create(t *testing.T) {
	category := models.Category{
		Name: "Test Category",
		Code: "001",
	}
	createdCategory, err := categoryRepo.Create(&category)
	assert.Nil(t, err)

	product := models.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Visible:     true,
		Images:      "test.jpg",
		CategoryID:  createdCategory.ID,
	}
	createdProduct, err := productRepo.Create(&product)
	assert.Nil(t, err)

	productVariation := models.ProductVariation{
		ProductID: createdProduct.ID,
		Name:      "Test Product Variation",
		Price:     59.99,
		Variation: "250g",
		Stock:     10,
	}
	productVariation2 := models.ProductVariation{
		ProductID: createdProduct.ID,
		Name:      "Test Product Variation",
		Price:     89.99,
		Variation: "500g",
		Stock:     10,
	}
	createdProductVariation, err := productVariationRepo.Create(&productVariation)
	assert.Nil(t, err)
	assert.Equal(t, productVariation.ProductID, createdProductVariation.ProductID)
	assert.Equal(t, productVariation.Name, createdProductVariation.Name)
	assert.Equal(t, productVariation.Price, createdProductVariation.Price)
	assert.Equal(t, productVariation.Variation, createdProductVariation.Variation)
	assert.Equal(t, productVariation.Stock, createdProductVariation.Stock)

	createdProductVariation2, err := productVariationRepo.Create(&productVariation2)
	assert.Nil(t, err)
	assert.Equal(t, productVariation2.ProductID, createdProductVariation2.ProductID)
	assert.Equal(t, productVariation2.Name, createdProductVariation2.Name)
	assert.Equal(t, productVariation2.Price, createdProductVariation2.Price)
	assert.Equal(t, productVariation2.Variation, createdProductVariation2.Variation)
	assert.Equal(t, productVariation2.Stock, createdProductVariation2.Stock)

	prd, err := productRepo.FindByID(createdProduct.ID)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(prd.ProductVariations))
	db.Delete(&createdProductVariation)
	db.Delete(&createdProductVariation2)
	db.Delete(&createdProduct)
	db.Delete(&createdCategory)
}


func TestProductVariationRepository_FindByProductID(t *testing.T) {
	category := models.Category{
		Name: "Test Category",
		Code: "001",
	}
	createdCategory, err := categoryRepo.Create(&category)
	assert.Nil(t, err)

	product := models.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Visible:     true,
		Images:      "test.jpg",
		CategoryID:  createdCategory.ID,
	}
	createdProduct, err := productRepo.Create(&product)
	assert.Nil(t, err)

	productVariation := models.ProductVariation{
		ProductID: createdProduct.ID,
		Name:      "Test Product Variation",
		Price:     59.99,
		Variation: "250g",
		Stock:     10,
	}
	productVariation2 := models.ProductVariation{
		ProductID: createdProduct.ID,
		Name:      "Test Product Variation",
		Price:     89.99,
		Variation: "500g",
		Stock:     10,
	}
	createdProductVariation, err := productVariationRepo.Create(&productVariation)
	assert.Nil(t, err)
	createdProductVariation2, err := productVariationRepo.Create(&productVariation2)
	assert.Nil(t, err)

	productVariations, err := productVariationRepo.ListByProductID(createdProduct.ID)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(productVariations))
	db.Delete(&createdProductVariation)
	db.Delete(&createdProductVariation2)
	db.Delete(&createdProduct)
	db.Delete(&createdCategory)
}

func TestProductVariationRepository_FindByID(t *testing.T) {
	category := models.Category{
		Name: "Test Category",
		Code: "001",
	}
	createdCategory, err := categoryRepo.Create(&category)
	assert.Nil(t, err)

	product := models.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Visible:     true,
		Images:      "test.jpg",
		CategoryID:  createdCategory.ID,
	}
	createdProduct, err := productRepo.Create(&product)
	assert.Nil(t, err)

	productVariation := models.ProductVariation{
		ProductID: createdProduct.ID,
		Name:      "Test Product Variation",
		Price:     59.99,
		Variation: "250g",
		Stock:     10,
	}
	createdProductVariation, err := productVariationRepo.Create(&productVariation)
	assert.Nil(t, err)

	productVariationFound, err := productVariationRepo.FindByID(createdProductVariation.ID)
	assert.Nil(t, err)
	assert.Equal(t, productVariationFound.ProductID, createdProductVariation.ProductID)
	assert.Equal(t, productVariationFound.Name, createdProductVariation.Name)
	assert.Equal(t, productVariationFound.Price, createdProductVariation.Price)
	assert.Equal(t, productVariationFound.Variation, createdProductVariation.Variation)
	assert.Equal(t, productVariationFound.Stock, createdProductVariation.Stock)
	db.Delete(&createdProductVariation)
	db.Delete(&createdProduct)
	db.Delete(&createdCategory)
}

func TestProductVariationRepository_Update(t *testing.T) {
	category := models.Category{
		Name: "Test Category",
		Code: "001",
	}
	createdCategory, err := categoryRepo.Create(&category)
	assert.Nil(t, err)

	product := models.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Visible:     true,
		Images:      "test.jpg",
		CategoryID:  createdCategory.ID,
	}
	createdProduct, err := productRepo.Create(&product)
	assert.Nil(t, err)

	productVariation := models.ProductVariation{
		ProductID: createdProduct.ID,
		Name:      "Test Product Variation",
		Price:     59.99,
		Variation: "250g",
		Stock:     10,
	}
	createdProductVariation, err := productVariationRepo.Create(&productVariation)
	assert.Nil(t, err)

	productVariationFound, err := productVariationRepo.FindByID(createdProductVariation.ID)
	assert.Nil(t, err)
	productVariationFound.Name = "Test Product Variation Updated"
	productVariationFound.Price = 69.99
	productVariationFound.Variation = "500g"
	productVariationFound.Stock = 20
	updatedProductVariation, err := productVariationRepo.Update(productVariationFound)
	assert.Nil(t, err)
	assert.Equal(t, productVariationFound.ProductID, updatedProductVariation.ProductID)
	assert.Equal(t, productVariationFound.Name, updatedProductVariation.Name)
	assert.Equal(t, productVariationFound.Price, updatedProductVariation.Price)
	assert.Equal(t, productVariationFound.Variation, updatedProductVariation.Variation)
	assert.Equal(t, productVariationFound.Stock, updatedProductVariation.Stock)
	db.Delete(&updatedProductVariation)
	db.Delete(&createdProduct)
	db.Delete(&createdCategory)
}