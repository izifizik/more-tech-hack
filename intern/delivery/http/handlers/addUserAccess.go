package handlers

import (
	"github.com/gin-gonic/gin"
	"more-tech-hack/intern/repository"
	"net/http"
)

func AddUserAccess(c *gin.Context) {
	jsonInput := struct {
		ModelId int    `json:"modelId"`
		UserId  string `json:"userId"`
	}{}

	err := c.ShouldBindJSON(&jsonInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid data",
		})
		return
	}

	err = repository.InsertUserAccess(jsonInput.UserId, jsonInput.ModelId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
