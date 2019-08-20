package handler

import "github.com/gin-gonic/gin"

func Root() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Root"})
	}
}
