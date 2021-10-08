package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Port string // port without ':'
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("can't load from file: " + err.Error())
	}
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8081"
	}
}
