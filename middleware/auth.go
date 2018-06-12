package middleware

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/stobita/golang-sns-sample/lib"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := string(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.JSON(400, lib.ErrorResponse("Auth header empty"))
			c.Abort()
			return
		}
		tokenString := strings.Split(authHeader, " ")[1]
		if result, userID := lib.TokenAuthenticate(tokenString); !result {
			c.JSON(401, lib.ErrorResponse("Invalid Token"))
			c.Abort()
			return
		} else {
			log.Println(userID)
			c.Set("userId", userID)
		}
		c.Next()
	}
}
