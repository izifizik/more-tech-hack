package app

import (
	"more-tech-hack/internal/config"
	"more-tech-hack/internal/delivery/http"
	"more-tech-hack/internal/services"
)

func Run() error {
	config.Load()
	config.ConnectDb()
	services.Init()
	err := http.Run()
	if err != nil {
		return err
	}
	return nil
}
