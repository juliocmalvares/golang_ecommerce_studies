package controllers

import (
	"ecomm/domain/services"
	"ecomm/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService services.ProductService
}

func InitProductController() *ProductController {
	service := services.InitProductService()
	if service == nil {
		return nil
	}
	return &ProductController{
		productService: *service,
	}
}

type IProductController interface {
	List() gin.HandlerFunc
	FindByName() gin.HandlerFunc
	FindByID() gin.HandlerFunc
	ListByCategoryID() gin.HandlerFunc
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
}

func (c *ProductController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := ProductCreateBody{}
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to bind request body",
				err,
			))
			ctx.Done()
			return
		}
		product, err := c.productService.Create(body.ToModel())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to save product",
				err,
			))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(
			http.StatusOK,
			product,
			"Success",
			nil,
		))
		ctx.Done()
	}
}

func (c *ProductController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body := ProductUpdateBody{}
		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to bind request body",
				err,
			))
			ctx.Done()
			return
		}
		product, err := c.productService.Update(body.ToModel())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to update product",
				err,
			))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(
			http.StatusOK,
			product,
			"Success",
			nil,
		))
		ctx.Done()
	}
}

func (c *ProductController) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := c.productService.List()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to fetch products",
				err,
			))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(
			http.StatusOK,
			products,
			"Success",
			nil,
		))
		ctx.Done()
	}
}

func (c *ProductController) FindByName() gin.HandlerFunc {
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
		product, err := c.productService.FindByName(name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to fetch product",
				err,
			))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(
			http.StatusOK,
			product,
			"Success",
			nil,
		))
		ctx.Done()
	}
}

func (c *ProductController) FindByID() gin.HandlerFunc {
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
		product, err := c.productService.FindByID(uint(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to fetch product",
				err,
			))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(
			http.StatusOK,
			product,
			"Success",
			nil,
		))
		ctx.Done()
	}
}

func (c *ProductController) ListByCategoryID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		stid, exists := ctx.GetQuery("category_id")
		id, err := strconv.Atoi(stid)
		if !exists {
			ctx.JSON(http.StatusBadRequest, utils.BuildDefaultResponse(
				http.StatusBadRequest,
				nil,
				"Category ID is required",
				nil,
			))
			ctx.Done()
			return
		}
		products, err := c.productService.ListByCategoryID(uint(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to fetch products",
				err,
			))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(
			http.StatusOK,
			products,
			"Success",
			nil,
		))
		ctx.Done()
	}
}

