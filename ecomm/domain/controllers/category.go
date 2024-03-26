package controllers

import (
	"ecomm/domain/services"
	"ecomm/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService services.CategoryService
}

func InitCategoryController() *CategoryController {
	service := services.InitCategoryService()
	if service == nil {
		return nil
	}
	return &CategoryController{
		categoryService: *service,
	}
}

type ICategoryController interface {
	List() gin.HandlerFunc
	FindByName() gin.HandlerFunc
	FindByID() gin.HandlerFunc
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
}

func (c *CategoryController) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categories, err := c.categoryService.List()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to fetch categories",
				err,
			))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(
			http.StatusOK,
			categories,
			"Success",
			nil,
		))
		ctx.Done()
	}
}

func (c *CategoryController) FindByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name, exists := ctx.GetQuery("name")
		if !exists {
			ctx.JSON(http.StatusBadRequest, utils.BuildDefaultResponse(
				http.StatusBadRequest,
				nil,
				"Name is required",
				nil,
			))
			ctx.Done()
			return
		}
		category, err := c.categoryService.FindByName(name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to fetch category",
				err,
			))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(
			http.StatusOK,
			category,
			"Success",
			nil,
		))
		ctx.Done()
	}
}

func (c *CategoryController) FindByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		stid, exists := ctx.GetQuery("id")
		id, err := strconv.Atoi(stid)
		if !exists {
			ctx.JSON(http.StatusBadRequest, utils.BuildDefaultResponse(
				http.StatusBadRequest,
				nil,
				"ID is required",
				nil,
			))
			ctx.Done()
			return
		}
		category, err := c.categoryService.FindByID(uint(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to fetch category",
				err,
			))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(
			http.StatusOK,
			category,
			"Success",
			nil,
		))
		ctx.Done()
	}
}

func (c *CategoryController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var categoryRequestBody CategoryCreateBody
		if err := ctx.BindJSON(&categoryRequestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildDefaultResponse(
				http.StatusBadRequest,
				nil,
				"Invalid request body",
				err,
			))
			ctx.Done()
			return
		}
		category, err := c.categoryService.Create(categoryRequestBody.ToModel())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to create category",
				err,
			))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(
			http.StatusOK,
			category,
			"Success",
			nil,
		))
		ctx.Done()
	}
}

func (c *CategoryController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var categoryRequestBody CategoryUpdateBody
		if err := ctx.BindJSON(&categoryRequestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildDefaultResponse(
				http.StatusBadRequest,
				nil,
				"Invalid request body",
				err,
			))
			ctx.Done()
			return
		}
		category, err := c.categoryService.Update(categoryRequestBody.ToModel())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to update category",
				err,
			))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(
			http.StatusOK,
			category,
			"Success",
			nil,
		))
		ctx.Done()
	}
}
