package handlers

import (
	"context"
	"github.com/Nerzal/gocloak/v9"
	"github.com/gin-gonic/gin"
	"more-tech-hack/internal/config"
	"net/http"
)


func Register(c *gin.Context) {
	jsonInput := struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		EMail     string `json:"e_mail"`
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
	token, err := client.LoginAdmin(ctx, "dima", "dimadima", "dima")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error with get admin auth " + err.Error(),
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

	userID, err := client.CreateUser(ctx, token.AccessToken, "dima", user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error with create user: " + err.Error(),
		})
		return
	}

	err = client.SetPassword(ctx, token.AccessToken, userID, "dima", jsonInput.Password, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":       "Error with set password: " + err.Error(),
		})
		return
	}

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
