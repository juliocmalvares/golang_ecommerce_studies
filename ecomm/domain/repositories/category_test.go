package repositories

import (
	"ecomm/domain/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategoryRepository_Create(t *testing.T) {
	category := models.Category{
		Name: "Test Category",
		Code: "001",
	}
	createdCategory, err := categoryRepo.Create(&category)
	assert.Nil(t, err)

	categoryChildren := models.Category{
		Name:     "Test Category Children",
		Code:     "002",
		ParentID: &createdCategory.ID,
	}
	createdCategoryChildren, err := categoryRepo.Create(&categoryChildren)
	assert.Nil(t, err)

	assert.Equal(t, category.Name, createdCategory.Name)
	assert.Equal(t, category.Code, createdCategory.Code)
	assert.Equal(t, createdCategory.ID, uint(1))
	assert.Equal(t, createdCategoryChildren.ID, uint(2))
	assert.Equal(t, *createdCategoryChildren.ParentID, createdCategory.ID)
	db.Delete(&createdCategoryChildren)
	db.Delete(&createdCategory)
}

func TestCategoryRepository_FindByID(t *testing.T) {
	category := models.Category{
		Name: "Test Category",
		Code: "001",
	}
	createdCategory, err := categoryRepo.Create(&category)
	assert.Nil(t, err)

	categoryChildren := models.Category{
		Name:     "Test Category Children",
		Code:     "002",
		ParentID: &createdCategory.ID,
	}
	createdCategoryChildren, err := categoryRepo.Create(&categoryChildren)
	assert.Nil(t, err)

	categoryFound, err := categoryRepo.FindByID(createdCategory.ID)
	assert.Nil(t, err)
	assert.Equal(t, categoryFound.Name, createdCategory.Name)
	assert.Equal(t, categoryFound.Code, createdCategory.Code)
	assert.Equal(t, categoryFound.ID, createdCategory.ID)
	assert.Equal(t, categoryFound.Childrens[0].Name, createdCategoryChildren.Name)
	assert.Equal(t, categoryFound.Childrens[0].Code, createdCategoryChildren.Code)
	db.Delete(&createdCategoryChildren)
	db.Delete(&createdCategory)
}

func TestCategoryRepository_FindByName(t *testing.T) {
	category := models.Category{
		Name: "Test Category",
		Code: "001",
	}
	createdCategory, err := categoryRepo.Create(&category)
	assert.Nil(t, err)

	categoryChildren := models.Category{
		Name:     "Test Category Children",
		Code:     "002",
		ParentID: &createdCategory.ID,
	}
	createdCategoryChildren, err := categoryRepo.Create(&categoryChildren)
	assert.Nil(t, err)

	categoryFound, err := categoryRepo.FindByName(createdCategory.Name)
	assert.Nil(t, err)
	assert.Equal(t, categoryFound.Name, createdCategory.Name)
	assert.Equal(t, categoryFound.Code, createdCategory.Code)
	assert.Equal(t, categoryFound.ID, createdCategory.ID)
	assert.Equal(t, categoryFound.Childrens[0].Name, createdCategoryChildren.Name)
	assert.Equal(t, categoryFound.Childrens[0].Code, createdCategoryChildren.Code)
	db.Delete(&createdCategoryChildren)
	db.Delete(&createdCategory)
}

func TestCategoryRepository_List(t *testing.T) {
	category := models.Category{
		Name: "Test Category",
		Code: "001",
	}
	createdCategory, err := categoryRepo.Create(&category)
	assert.Nil(t, err)

	categoryChildren := models.Category{
		Name:     "Test Category Children",
		Code:     "002",
		ParentID: &createdCategory.ID,
	}
	createdCategoryChildren, err := categoryRepo.Create(&categoryChildren)
	assert.Nil(t, err)

	categories, err := categoryRepo.List()
	assert.Nil(t, err)
	assert.Equal(t, len(categories), 2)
	assert.Equal(t, categories[0].Name, createdCategory.Name)
	assert.Equal(t, categories[0].Code, createdCategory.Code)
	assert.Equal(t, categories[0].ID, createdCategory.ID)
	assert.Equal(t, categories[1].Name, createdCategoryChildren.Name)
	assert.Equal(t, categories[1].Code, createdCategoryChildren.Code)
	db.Delete(&createdCategoryChildren)
	db.Delete(&createdCategory)
	categories, err = categoryRepo.List()
	assert.Nil(t, err)
	assert.Equal(t, len(categories), 0)
}

func TestCategoryRepository_Update(t *testing.T) {
	category := models.Category{
		Name: "Test Category",
		Code: "001",
	}
	createdCategory, err := categoryRepo.Create(&category)
	assert.Nil(t, err)
	up := models.Category{
		ID:       createdCategory.ID,
		Name:     "Test Category Updated",
		Code:     "002",
		ParentID: nil,
	}
	updatedCategory, err := categoryRepo.Update(&up)
	assert.Nil(t, err)
	assert.NotEqual(t, updatedCategory.Name, createdCategory.Name)
	assert.Equal(t, updatedCategory.Name, "Test Category Updated")
	assert.Equal(t, updatedCategory.Code, "002")
	assert.Equal(t, updatedCategory.ID, createdCategory.ID)
	db.Delete(&updatedCategory)
}
