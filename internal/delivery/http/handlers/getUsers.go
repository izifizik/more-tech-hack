package handlers

import (
	"github.com/gin-gonic/gin"
	"more-tech-hack/internal/repository"
	"net/http"
	"strconv"
)

func GetUsers(c *gin.Context) {
	modelId, err := strconv.Atoi(c.Param("modelId"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "invalid modelId",
		})
	}
	users, err := repository.GetUsersByModelId(modelId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
