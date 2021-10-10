package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"more-tech-hack/intern/repository"
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
	value, _ := c.Get("userId")
	b := repository.UpdateBalance(value.(string))
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"models":  model,
		"balance": b,
	})
}

func GetModel(c *gin.Context) {
	i, _ := strconv.Atoi(c.Param("id"))
	value, _ := c.Get("userId")
	repository.UpdateBalance(value.(string))
	model, err := repository.GetModel(i)
	b := repository.UpdateBalance(value.(string))
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"model":   model,
		"balance": b,
	})
}
