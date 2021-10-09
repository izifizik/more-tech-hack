package handlers

import (
	"github.com/gin-gonic/gin"
	"more-tech-hack/internal/services"
	"net/http"
)

func GetDataset(c *gin.Context) {
	urn := c.DefaultQuery("urn", "urn:li:dataset:(urn:li:dataPlatform:hive,SampleHiveDataset,PROD)")
	resp := services.GetDataset(urn)
	c.JSON(http.StatusOK, gin.H{
		"dataset": resp.Dataset,
	})
}

func Browse(c *gin.Context) {
	path := c.DefaultQuery("path", "")
	resp := services.Browse(path)
	c.JSON(http.StatusOK, gin.H{
		"groups": resp.Browse,
	})
}
