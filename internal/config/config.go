package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Port             string // port without ':'
	KeyAdminUsername string
	KeyAdminPassword string
	KeyHttpPath      string
	KeyRealm         string
	KeySecret        string
	KeyClisentId     string
	CookieDataHub    string
	DataHubUrl       string
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
	KeyAdminUsername = os.Getenv("KC_ADMIN_USERNAME")
	KeyAdminPassword = os.Getenv("KC_ADMIN_PASSWORD")
	KeyHttpPath = os.Getenv("KC_CLIENT_PATH")
	KeyRealm = os.Getenv("KC_REALM")
	KeySecret = os.Getenv("KC_SECRET")
	KeyClisentId = os.Getenv("KC_CLIENT")
	CookieDataHub = os.Getenv("COOKIE_DATAHUB")
	DataHubUrl = os.Getenv("DATAHUB_URL")
}
