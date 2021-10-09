package handlers

import (
	"context"
	"github.com/Nerzal/gocloak/v9"
	"github.com/gin-gonic/gin"
	"log"
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
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Error with login: " + err.Error(),
		})
		return
	}

	params := gocloak.GetUsersParams{
		Username: gocloak.StringP(jsonInput.Username),
	}

	user, err := config.Client.GetUsers(ctx, config.AdminToken.AccessToken, config.KeyRealm, params)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}

	if len(user) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "ok",
		"access_token":  login.AccessToken,
		"refresh_token": login.RefreshToken,
		"exp_access":    login.ExpiresIn,
		"exp_refresh":   login.RefreshExpiresIn,
		"user":          user[0],
	})
}
