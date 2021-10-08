package app

import (
	"more-tech-hack/internal/config"
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
