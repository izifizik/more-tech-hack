package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"more-tech-hack/internal/repository"
	"net/http"
	"strconv"
)

func GetModels(c *gin.Context) {
	userId, isExist := c.Get("userId")
	if !isExist {
		fmt.Println("ошибка")
	}
	a, isOk := userId.(string)
	if !isOk {
		log.Println("ошибка")
	}
	model, err := repository.GetModels(a)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"models": model,
	})
}

func GetModel(c *gin.Context) {
	i, _ := strconv.Atoi(c.Param("id"))
	model, err := repository.GetModel(i)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"model": model,
	})
}
