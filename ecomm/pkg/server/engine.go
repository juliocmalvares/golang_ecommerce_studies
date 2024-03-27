package engine

import (
	"ecomm/domain/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Server struct {
	Engine *gin.Engine
}

func New() *Server {
	router := Server{Engine: gin.New()}
	router.middlewares()
	router.routes()
	if viper.GetString("ENV") == "dev" {
		gin.SetMode(gin.DebugMode)
	} else if viper.GetString("ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	return &router
}

func (r *Server) middlewares() {
	r.Engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Engine.Use(gin.Logger())
	r.Engine.Use(gin.Recovery())
}

func (r *Server) routes() {

	apiGroup := r.Engine.Group("/api")
	{
		//User
		userGroup := apiGroup.Group("/user")
		{
			userCtr := controllers.InitUserController()
			userGroup.POST("/create", userCtr.Create())
			userGroup.GET("/find", userCtr.FindByID())
			userGroup.PUT("/update", userCtr.Update())
		}
		// Category
		categoryGroup := apiGroup.Group("/category")
		{
			categoryCtr := controllers.InitCategoryController()
			categoryGroup.GET("/list", categoryCtr.List())
			// categoryGroup.GET("/find", categoryCtr.FindByID())
			categoryGroup.GET("/find", categoryCtr.FindByName())
			categoryGroup.POST("/create", categoryCtr.Create())
			categoryGroup.PUT("/update", categoryCtr.Update())

		}

		// //Product
		productGroup := apiGroup.Group("/product")
		{
			productCtr := controllers.InitProductController()
			productGroup.POST("/create", productCtr.Create())
			productGroup.PUT("/update", productCtr.Update())
			productGroup.GET("/list", productCtr.List())
			productGroup.GET("/find", productCtr.FindByID())
			productGroup.GET("/findByName", productCtr.FindByName())
			productGroup.GET("/listByCategory", productCtr.ListByCategoryID())
		}
	}
}
