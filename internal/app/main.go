package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"more-tech-hack/internal/app/config"
)

func Run() error {
	config.Load()

	app := gin.Default()
	gin.SetMode(gin.DebugMode)

	err := app.Run(config.Port)
	if err != nil {
		log.Println("Internal server err")
	}
	return nil
}
