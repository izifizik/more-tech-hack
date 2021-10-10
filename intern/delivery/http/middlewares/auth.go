package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"more-tech-hack/intern/config"
	"net/http"
)

func AuthMiddleware(c *gin.Context) {
	_, claims, err := config.Client.DecodeAccessToken(context.Background(), c.GetHeader("Authorization"), config.KeyRealm, "")
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid token " + err.Error(),
		})
		return
	}
	a := *claims
	c.Set("userId", a["sub"])
	c.Next()
}
