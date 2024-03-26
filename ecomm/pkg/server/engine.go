package engine

import (
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
		apiGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
