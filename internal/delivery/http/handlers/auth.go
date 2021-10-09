package handlers

import (
	"context"
	"github.com/Nerzal/gocloak/v9"
	"github.com/gin-gonic/gin"
	"more-tech-hack/internal/config"
	"net/http"
)

func Auth(c *gin.Context) {
	ctx := context.Background()

	jsonInput := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.ShouldBindJSON(&jsonInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error with pars json: " + err.Error(),
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

	params := gocloak.GetUsersParams{
		Username: gocloak.StringP(jsonInput.Username),
	}

	user, err := config.Client.GetUsers(ctx, config.AdminToken.AccessToken, config.KeyRealm, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error with get client: " + err.Error(),
		})
		return
	}

	if len(user) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "OK",
		"access_token":  login.AccessToken,
		"refresh_token": login.RefreshToken,
		"exp_access":    login.ExpiresIn,
		"exp_refresh":   login.RefreshExpiresIn,
		"user_access": user,
	})
}
