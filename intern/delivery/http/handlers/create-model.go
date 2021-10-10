package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"more-tech-hack/intern/model"
	"more-tech-hack/intern/repository"
	"net/http"
)

func CreateModel(c *gin.Context) {
	var m model.Model
	err := c.ShouldBindJSON(&m)
	if err != nil {
		log.Println(err)
	}
	err = repository.InsertModel(&m)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
