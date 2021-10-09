package handlers

import (
	"github.com/gin-gonic/gin"
	"more-tech-hack/internal/services"
	"net/http"
)

func GetDataset(c *gin.Context) {
	resp := services.GetDataset()
	c.JSON(http.StatusOK, gin.H{
		"fields": resp.Dataset.SchemaMetadata.Fields,
	})
}

func Browse(c *gin.Context) {
	resp := services.Browse()
	c.JSON(http.StatusOK, gin.H{
		"fields": resp,
	})
}
