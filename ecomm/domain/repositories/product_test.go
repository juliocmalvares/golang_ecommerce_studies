package repositories

import (
	"ecomm/domain/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductRepository_Create(t *testing.T) {
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

	assert.Equal(t, product.Name, createdProduct.Name)
	assert.Equal(t, product.Description, createdProduct.Description)
	assert.Equal(t, product.Visible, createdProduct.Visible)
	assert.Equal(t, product.Images, createdProduct.Images)
	assert.Equal(t, product.CategoryID, createdProduct.CategoryID)
	db.Delete(&createdProduct)
	db.Delete(&createdCategory)
}

func TestProductRepository_FindByID(t *testing.T) {
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

	productFound, err := productRepo.FindByID(createdProduct.ID)
	assert.Nil(t, err)
	assert.Equal(t, productFound.Name, createdProduct.Name)
	assert.Equal(t, productFound.Description, createdProduct.Description)
	assert.Equal(t, productFound.Visible, createdProduct.Visible)
	assert.Equal(t, productFound.Images, createdProduct.Images)
	assert.Equal(t, productFound.CategoryID, createdProduct.CategoryID)
	assert.Equal(t, productFound.ID, createdProduct.ID)
	db.Delete(&createdProduct)
	db.Delete(&createdCategory)
}

func TestProductRepository_List(t *testing.T) {
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
	_, err = productRepo.Create(&product)
	assert.Nil(t, err)

	products, err := productRepo.List()
	assert.Nil(t, err)
	assert.Equal(t, len(products), 1)
	db.Delete(&products[0])
	db.Delete(&createdCategory)
}

func TestProductRepository_ListByCategoryID(t *testing.T) {
	category := models.Category{
		Name: "Test Category",
		Code: "001",
	}
	createdCategory, err := categoryRepo.Create(&category)
	assert.Nil(t, err)

	category2 := models.Category{
		Name: "Test Category 2",
		Code: "002",
	}
	createdCategory2, err := categoryRepo.Create(&category2)
	assert.Nil(t, err)

	product := models.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Visible:     true,
		Images:      "test.jpg",
		CategoryID:  createdCategory.ID,
	}
	_, err = productRepo.Create(&product)
	assert.Nil(t, err)

	product2 := models.Product{
		Name:        "Test Product 2",
		Description: "Test Description 2",
		Visible:     true,
		Images:      "test2.jpg",
		CategoryID:  createdCategory2.ID,
	}
	_, err = productRepo.Create(&product2)
	assert.Nil(t, err)

	products, err := productRepo.ListByCategoryID(createdCategory.ID)
	assert.Nil(t, err)
	assert.Equal(t, len(products), 1)
	assert.Equal(t, products[0].CategoryID, createdCategory.ID)
	assert.Equal(t, products[0].Name, product.Name)
	db.Delete(&product)
	db.Delete(&product2)
	db.Delete(&createdCategory)
	db.Delete(&createdCategory2)
}

func TestProductRepository_Update(t *testing.T) {
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

	createdProduct.Name = "Test Product Updated"
	updatedProduct, err := productRepo.Update(createdProduct)
	assert.Nil(t, err)
	assert.NotEqual(t, updatedProduct.Name, "Test Product")
	assert.Equal(t, updatedProduct.Name, "Test Product Updated")
	assert.Equal(t, updatedProduct.Description, createdProduct.Description)
	assert.Equal(t, updatedProduct.Visible, createdProduct.Visible)
	assert.Equal(t, updatedProduct.Images, createdProduct.Images)
	assert.Equal(t, updatedProduct.CategoryID, createdProduct.CategoryID)
	assert.Equal(t, updatedProduct.ID, createdProduct.ID)
	db.Delete(&updatedProduct)
	db.Delete(&createdCategory)
}
