package http

import (
	"more-tech-hack/internal/config"

	"github.com/gin-gonic/gin"
)

func Run() error {
	app := gin.Default()
	gin.SetMode(gin.DebugMode)

	err := app.Run(config.Port)
	return err
}
