package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(c *gin.Context) {



	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
	})
}