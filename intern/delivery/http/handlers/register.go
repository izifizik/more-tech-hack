package handlers

import (
	"context"
	"github.com/Nerzal/gocloak/v9"
	"github.com/gin-gonic/gin"
	"more-tech-hack/intern/config"
	"more-tech-hack/intern/repository"
	"net/http"
)

func Register(c *gin.Context) {
	ctx := context.Background()
	jsonInput := struct {
		FirstName string `json:"userFirstName"`
		LastName  string `json:"userLastName"`
		EMail     string `json:"userEmail"`
		Username  string `json:"username"`
		Password  string `json:"password"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error with pars json: " + err.Error(),
		})
		return
	}

	user := gocloak.User{
		FirstName: gocloak.StringP(jsonInput.FirstName),
		LastName:  gocloak.StringP(jsonInput.LastName),
		Email:     gocloak.StringP(jsonInput.EMail),
		Enabled:   gocloak.BoolP(true),
		Username:  gocloak.StringP(jsonInput.Username),
	}

	userID, err := config.Client.CreateUser(ctx, config.AdminToken.AccessToken, config.KeyRealm, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error with create user: " + err.Error(),
		})
		return
	}

	err = repository.InsertUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error with database insert: " + err.Error(),
		})
		return
	}

	err = config.Client.SetPassword(ctx, config.AdminToken.AccessToken, userID, config.KeyRealm, jsonInput.Password, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error with set password: " + err.Error(),
		})
		return
	}

	login, err := config.Client.Login(ctx, config.KeyClientId, config.KeySecret, config.KeyRealm, jsonInput.Username, jsonInput.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error with login: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "ok",
		"userFirstName": jsonInput.FirstName,
		"userLastLame":  jsonInput.LastName,
		"userEmail":     jsonInput.EMail,
		"username":      jsonInput.Username,
		"token":         login,
		"balance":       0,
	})
}
