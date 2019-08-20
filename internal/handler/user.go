package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stobita/golang-sns-sample/internal/lib"
	"github.com/stobita/golang-sns-sample/internal/model"
)

type UserPutJSON struct {
	Email string `json:"email"`
}

func UserUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json UserPutJSON
		if c.BindJSON(&json) != nil {
			c.JSON(400, lib.ErrorResponse("Invalid Params"))
			return
		}
		email := json.Email
		userID := c.Keys["userId"].(int64)
		err := model.NewUser(email).Update(userID)
		if err != nil {
			c.JSON(400, lib.ErrorResponse(err.Error()))
			return
		}
		c.AbortWithStatus(200)
	}
}

func GetUserPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Keys["userId"].(int64)
		result := model.User{ID: userID}.GetPosts()
		c.JSON(200, result)
	}
}
