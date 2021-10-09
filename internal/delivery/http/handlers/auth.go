package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(c *gin.Context) {

	josnInput := struct {
		jsonInput := struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		EMail string `json:"e_mail"`
		Username string `json:"username"`
	}{}
	}{}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
	})
}