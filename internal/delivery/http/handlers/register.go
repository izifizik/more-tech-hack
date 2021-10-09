package handlers

import (
	"github.com/Nerzal/gocloak"
	"github.com/gin-gonic/gin"
	"log"
	"more-tech-hack/internal/config"
	"net/http"
)

func Register(c *gin.Context) {
	config.LoadKC()

	client := gocloak.NewClient(config.KeyHttpPath)
	AdminToken, err := client.LoginAdmin("andrey", "andrey", "demo")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error with get admin auth " + err.Error(),
		})
		return
	}

	jsonInput := struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		EMail string `json:"e_mail"`
		Username string `json:"username"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "not all parameters are specified",
		})
		return
	}

	//user := gocloak.User{
	//	FirstName: jsonInput.FirstName,
	//	LastName: jsonInput.LastName,
	//	Email: jsonInput.EMail,
	//	Enabled: true,
	//	Username: jsonInput.Username,
	//}
	log.Println(AdminToken.AccessToken)
	createUser, err := client.CreateUser(AdminToken.AccessToken, "demo", gocloak.User{})
	log.Println(createUser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok",
			"access_token": "",
			"refresh_token": AdminToken.RefreshToken,
			"exp": AdminToken.ExpiresIn,
		})
	}

}
