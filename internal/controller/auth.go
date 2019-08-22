package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/stobita/golang-sns-sample/internal/lib"
	"github.com/stobita/golang-sns-sample/internal/model"
	"github.com/stobita/golang-sns-sample/internal/presenter"
)

type signUpRequestBody struct {
	Email                string `json:"email" binding:"required"`
	Password             string `json:"password" binding:"required"`
	ConfirmationPassword string `json:"confirmationPassword" binding:"required"`
}

type signInRequestBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (c *controller) SignUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var json signUpRequestBody
		if err := ctx.ShouldBindJSON(&json); err != nil {
			log.Error(err)
			ctx.JSON(http.StatusBadRequest, presenter.ErrorResponse("Invalid params"))
			return
		}
		email := json.Email
		password := json.Password
		confirmationPassword := json.ConfirmationPassword
		if password != confirmationPassword {
			ctx.JSON(http.StatusBadRequest, presenter.ErrorResponse("Password Mismatch"))
			return
		}
		encryptedPassword, err := lib.GetEncryptedPassword(password)
		if err != nil || encryptedPassword == "" {
			log.Error(err)
			ctx.JSON(http.StatusInternalServerError, presenter.ErrorResponse("encrypt error"))
			return
		}
		model := &model.User{
			Email:    email,
			Password: encryptedPassword,
		}
		if err := c.repository.CreateUser(model); err != nil {
			log.Error(err)
			ctx.JSON(http.StatusInternalServerError, presenter.ErrorResponse("Failed create user"))
			return

		}
		ctx.AbortWithStatus(http.StatusOK)
	}
}

func (c *controller) SignIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var json signInRequestBody
		if err := ctx.ShouldBindJSON(&json); err != nil {
			log.Error(err)
			ctx.JSON(http.StatusBadRequest, presenter.ErrorResponse("Invalid params"))
			return
		}
		email := json.Email
		password := json.Password
		user, err := c.repository.FindUserByEmail(email)
		if err != nil || user == nil {
			log.Error(err)
			ctx.JSON(http.StatusForbidden, presenter.ErrorResponse("User not found"))
			return
		}
		if lib.ComparePassword(password, user.Password) {
			if tokenString, err := lib.GenerateTokenString(user.ID); err == nil {
				ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
				return
			} else {
				log.Error(err)
				ctx.JSON(http.StatusInternalServerError, presenter.ErrorResponse("token generate error"))
				return
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, presenter.ErrorResponse("Invalid email or password"))
			return
		}
	}
}
