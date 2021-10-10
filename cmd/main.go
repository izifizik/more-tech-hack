package main

import (
	"log"
	"more-tech-hack/intern/app"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
