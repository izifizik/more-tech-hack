package app

import (
	"more-tech-hack/internal/app/config"
	"more-tech-hack/internal/delivery/http"
)

func Run() error {
	config.Load()
	err := http.Run()
	if err != nil {
		return err
	}
	return nil
}
