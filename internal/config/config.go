package config

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Nerzal/gocloak/v9"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	Port             string // port without ':'
	KeyAdminUsername string
	KeyAdminPassword string
	KeyHttpPath      string
	KeyRealm         string
	KeySecret        string
	KeyClientId      string
	CookieDataHub    string
	DataHubUrl       string
	PortDatabase     string
	HostDb           string
	UserDb           string
	PasswdDb         string
	NameDb           string
	Client           gocloak.GoCloak
	AdminToken       *gocloak.JWT
	Db               *sql.DB
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
	KeyClientId = os.Getenv("KC_CLIENT")
	CookieDataHub = os.Getenv("COOKIE_DATAHUB")
	DataHubUrl = os.Getenv("DATAHUB_URL")
	PortDatabase = os.Getenv("PORT_DB")
	HostDb = os.Getenv("HOST_DB")
	UserDb = os.Getenv("USER_DB")
	PasswdDb = os.Getenv("PASSWD_DB")
	NameDb = os.Getenv("NAME_DB")

	Client = gocloak.NewClient(KeyHttpPath)
	ctx := context.Background()
	AdminToken, err = Client.LoginAdmin(ctx, KeyAdminUsername, KeyAdminPassword, KeyRealm)
	if err != nil {
		log.Println("Get keycloak error")
	}
}

func ConnectDb() {
	PortDb, err := strconv.Atoi(PortDatabase)
	if err != nil {
		log.Fatal(err)
	}
	// connection string
	psqlconnect := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HostDb, PortDb, UserDb, PasswdDb, NameDb)

	// open database
	Db, err = sql.Open("postgres", psqlconnect)
	if err != nil {
		log.Println("Cannot open database")
		return
	}

	fmt.Println("Connected to Postgresql!")
}
