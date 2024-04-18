package controllers

import (
	"ecomm/domain/services"
	"ecomm/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderService     services.IOrderService
	OrderItemService services.IOrderItemService
}

func InitOrderController() *OrderController {
	OrderService := services.InitOrderService()
	OrderItemService := services.InitOrderItemService()

	if OrderService == nil || OrderItemService == nil {
		return nil
	}
	return &OrderController{
		OrderService:     OrderService,
		OrderItemService: OrderItemService,
	}
}

type IOrderController interface {
	List() gin.HandlerFunc
	ListByUserID() gin.HandlerFunc
	FindByID() gin.HandlerFunc
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
	AddOrderItem() gin.HandlerFunc
	RemoveOrderItem() gin.HandlerFunc
}

func (c *OrderController) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		orders, err := c.OrderService.List()
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				utils.BuildDefaultResponse(
					http.StatusInternalServerError,
					nil,
					"Failed to bind request body",
					err,
				))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(http.StatusOK, orders, "Success", nil))
		ctx.Done()
	}
}

func (c *OrderController) ListByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		stid, exists := ctx.GetQuery("user_id")
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
		orders, err := c.OrderService.ListByUserID(uint(id))
		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				utils.BuildDefaultResponse(
					http.StatusInternalServerError,
					nil,
					"Failed to bind request body",
					err,
				))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(http.StatusOK, orders, "Success", nil))
		ctx.Done()
	}
}

func (c *OrderController) FindByID() gin.HandlerFunc {
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
		order, err := c.OrderService.FindByID(uint(id))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to fetch order",
				err,
			))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(
			http.StatusOK,
			order,
			"Success",
			nil,
		))
		ctx.Done()
	}
}

func (c *OrderController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var order OrderCreateBody
		if err := ctx.ShouldBindJSON(&order); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildDefaultResponse(
				http.StatusBadRequest,
				nil,
				"Failed to bind request body",
				err,
			))
			ctx.Done()
			return
		}
		createdOrder, err := c.OrderService.Create(order.ToModel())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to create order",
				err,
			))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(
			http.StatusOK,
			createdOrder,
			"Success",
			nil,
		))
		ctx.Done()
	}
}

// func (c *OrderController) Update() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var order OrderUpdateBody
// 		if err := ctx.ShouldBindJSON(&order); err != nil {
// 			ctx.JSON(http.StatusBadRequest, utils.BuildDefaultResponse(
// 				http.StatusBadRequest,
// 				nil,
// 				"Failed to bind request body",
// 				err,
// 			))
// 			ctx.Done()
// 			return
// 		}
// 		updatedOrder, err := c.OrderService.Update(order.ToModel())
// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
// 				http.StatusInternalServerError,
// 				nil,
// 				"Failed to update order",
// 				err,
// 			))
// 			ctx.Done()
// 			return
// 		}
// 		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(
// 			http.StatusOK,
// 			updatedOrder,
// 			"Success",
// 			nil,
// 		))
// 		ctx.Done()
// 	}
// }

func (c *OrderController) AddOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var orderItem OrderItemAddBody
		if err := ctx.ShouldBindJSON(&orderItem); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildDefaultResponse(
				http.StatusBadRequest,
				nil,
				"Failed to bind request body",
				err,
			))
			ctx.Done()
			return
		}
		createdOrderItem, err := c.OrderItemService.AddOrderItem(orderItem.ToModel())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to create order item",
				err,
			))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(
			http.StatusOK,
			createdOrderItem,
			"Success",
			nil,
		))
		ctx.Done()
	}
}

func (c *OrderController) RemoveOrderItem() gin.HandlerFunc {
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
		orderItem, err := c.OrderItemService.FindByID(uint(id))
		err = c.OrderItemService.RemoveOrderItem(orderItem)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to remove order item",
				err,
			))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(
			http.StatusOK,
			nil,
			"Success",
			nil,
		))
		ctx.Done()
	}
}
