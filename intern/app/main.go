package app

import (
	"more-tech-hack/intern/config"
	"more-tech-hack/intern/delivery/http"
	"more-tech-hack/intern/services"
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
