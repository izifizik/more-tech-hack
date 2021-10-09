package handlers

import (
	"context"
	"github.com/Nerzal/gocloak/v9"
	"github.com/gin-gonic/gin"
	"more-tech-hack/internal/config"
	"net/http"
)

func Auth(c *gin.Context) {

	jsonInput := struct {
		Username  string `json:"username"`
		Password  string `json:"password"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error with pars json: " + err.Error(),
		})
		return
	}

	client := gocloak.NewClient(config.KeyHttpPath)
	ctx := context.Background()

	login, err := client.Login(ctx, config.KeyClisentId, config.KeySecret, "dima", jsonInput.Username, jsonInput.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error with login: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"access_token": login.AccessToken,
		"refresh_token": login.RefreshToken,
		"exp_access": login.ExpiresIn,
		"exp_refresh": login.RefreshExpiresIn,
	})
}
