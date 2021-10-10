package handlers

import (
	"github.com/gin-gonic/gin"
	"more-tech-hack/intern/repository"
	"more-tech-hack/intern/services"
	"net/http"
)

func GetDataset(c *gin.Context) {
	urn := c.DefaultQuery("urn", "urn:li:dataset:(urn:li:dataPlatform:hive,SampleHiveDataset,PROD)")
	resp := services.GetDataset(urn)
	value, _ := c.Get("userId")
	repository.UpdateBalance(value.(string))
	c.JSON(http.StatusOK, gin.H{
		"dataset": resp.Dataset,
	})
}

func Browse(c *gin.Context) {
	path := c.Param("path")
	resp := services.Browse(path)
	c.JSON(http.StatusOK, gin.H{
		"browse": resp.Browse,
	})
}
