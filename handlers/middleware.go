package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
)

const userIdHeader = "X-Tantum-Authorization"

func UserIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader(userIdHeader)
		if header == "" {
			log.Println("UserID not provided")
			return
		}
		c.Next()
	}
}
