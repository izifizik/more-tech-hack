package http

import (
	"more-tech-hack/internal/config"
	"more-tech-hack/internal/delivery/http/handlers"
	"more-tech-hack/internal/delivery/http/middlewares"

	"github.com/gin-gonic/gin"
)

func Run() error {
	app := gin.Default()
	gin.SetMode(gin.DebugMode)

	app.POST("/reg", handlers.Register)
	app.POST("/auth", handlers.Auth)

	app.Use(middlewares.AuthMiddleware)
	app.GET("/dataset", handlers.GetDataset)
	app.GET("/browse", handlers.Browse)

	err := app.Run("0.0.0.0:" + config.Port)

	return err
}
