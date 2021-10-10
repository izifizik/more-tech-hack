package handlers

import (
	"context"
	"github.com/Nerzal/gocloak/v9"
	"github.com/gin-gonic/gin"
	"log"
	"more-tech-hack/intern/config"
	"more-tech-hack/intern/repository"
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

	m, err := repository.GetUser(*user[0].ID)
	if err != nil {
		log.Println()
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "ok",
		"accessToken":   login.AccessToken,
		"refreshToken":  login.RefreshToken,
		"expAccess":     login.ExpiresIn,
		"expRefresh":    login.RefreshExpiresIn,
		"userFirstName": user[0].FirstName,
		"userLastName":  user[0].LastName,
		"userEmail":     user[0].Email,
		"balance":       m.Balance,
		"userId":        m.UserId,
	})
}
