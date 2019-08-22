package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/stobita/golang-sns-sample/internal/lib"
	"github.com/stobita/golang-sns-sample/internal/presenter"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := string(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.JSON(http.StatusForbidden, presenter.ErrorResponse("Auth header empty"))
			c.Abort()
			return
		}
		tokenString := strings.Split(authHeader, " ")[1]
		if result, userID := lib.TokenAuthenticate(tokenString); !result {
			c.JSON(http.StatusForbidden, presenter.ErrorResponse("Invalid Token"))
			c.Abort()
			return
		} else {
			c.Set("userId", userID)
		}
		c.Next()
	}
}
