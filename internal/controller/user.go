package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/stobita/golang-sns-sample/internal/presenter"
)

type updateUserRequestBody struct {
	Email string `json:"email"`
}

func (c *controller) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var json updateUserRequestBody
		if err := ctx.ShouldBindJSON(&json); err != nil {
			log.Error(err)
			ctx.JSON(http.StatusBadRequest, presenter.ErrorResponse("Invalid params"))
			return
		}
		email := json.Email
		userID := ctx.Keys["userId"].(uint)

		user, err := c.repository.FindUser(userID)
		if err != nil {
			log.Error(err)
			ctx.JSON(http.StatusBadRequest, presenter.ErrorResponse("User not found"))
			return
		}

		user.Email = email

		if err := c.repository.UpdateUser(user); err != nil {
			log.Error(err)
			ctx.JSON(http.StatusInternalServerError, presenter.ErrorResponse("Failed update user"))
			return
		}

		ctx.AbortWithStatus(http.StatusOK)
	}
}

func (c *controller) GetUserPosts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.Keys["userId"].(uint)
		posts, err := c.repository.GetUserPosts(userID)
		if err != nil {
			log.Error(err)
			ctx.JSON(http.StatusInternalServerError, presenter.ErrorResponse("Failed get user post"))
			return
		}
		json, err := presenter.PostsResponse(posts)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, presenter.ErrorResponse("response data error"))
			return
		}

		ctx.JSON(200, json)
	}
}
