package main

import (
	"log"
	"more-tech-hack/internal/app"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
