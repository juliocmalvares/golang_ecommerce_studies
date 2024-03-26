package controllers

import (
	"ecomm/domain/services"
	"ecomm/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func InitUserController() *UserController {
	service := services.InitUserService()
	if service == nil {
		return nil
	}
	return &UserController{
		userService: *service,
	}
}

type IUserController interface {
	FindByID() gin.HandlerFunc
	Create() gin.HandlerFunc
	Update() gin.HandlerFunc
}

func (c *UserController) FindByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		stid, _ := ctx.GetQuery("id")
		id, err := strconv.Atoi(stid)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildDefaultResponse(
				http.StatusBadRequest,
				nil,
				"Invalid ID",
				err,
			))
			ctx.Done()
			return
		}
		user, err := c.userService.FindByID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, utils.BuildDefaultResponse(
				http.StatusNotFound,
				nil,
				"User not found",
				err,
			))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(
			http.StatusOK,
			user,
			"User found",
			nil,
		))
		ctx.Done()
	}
}

func (c *UserController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userRequestBody UserCreateBody
		if err := ctx.BindJSON(&userRequestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildDefaultResponse(
				http.StatusBadRequest,
				nil,
				"Invalid request body",
				err,
			))
			ctx.Done()
			return
		}
		user, err := c.userService.Create(userRequestBody.ToModel())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to create user",
				err,
			))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusCreated, utils.BuildDefaultResponse(
			http.StatusCreated,
			user,
			"User created",
			nil,
		))
		ctx.Done()
	}
}

func (c *UserController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userRequestBody UserUpdateBody
		if err := ctx.BindJSON(&userRequestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, utils.BuildDefaultResponse(
				http.StatusBadRequest,
				nil,
				"Invalid request body",
				err,
			))
			ctx.Done()
			return
		}
		user, err := c.userService.Update(userRequestBody.ToModel())
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.BuildDefaultResponse(
				http.StatusInternalServerError,
				nil,
				"Failed to update user",
				err,
			))
			ctx.Done()
			return
		}
		ctx.JSON(http.StatusOK, utils.BuildDefaultResponse(
			http.StatusOK,
			user,
			"User updated",
			nil,
		))
		ctx.Done()
	}
}
