package http

import (
	"github.com/gin-contrib/cors"
	"more-tech-hack/internal/config"
	"more-tech-hack/internal/delivery/http/handlers"
	"more-tech-hack/internal/delivery/http/middlewares"

	"github.com/gin-gonic/gin"
)

func Run() error {
	app := gin.Default()
	corsMy := cors.DefaultConfig()
	corsMy.AllowHeaders = append(corsMy.AllowHeaders, "Authorization", "Access-Control-Allow-Origin")
	corsMy.AllowAllOrigins = true
	corsMy.AllowCredentials = true
	corsMy.AllowBrowserExtensions = true
	app.Use(cors.New(corsMy))
	gin.SetMode(gin.DebugMode)

	app.POST("/reg", handlers.Register)
	app.POST("/auth", handlers.Auth)

	app.Use(middlewares.AuthMiddleware)
	app.GET("/dataset", handlers.GetDataset)
	app.GET("/browse/:path", handlers.Browse)
	app.GET("/browse", handlers.Browse)

	err := app.Run("0.0.0.0:" + config.Port)

	return err
}
