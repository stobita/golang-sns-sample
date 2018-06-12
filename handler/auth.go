package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stobita/golang-sns-sample/lib"
	"github.com/stobita/golang-sns-sample/model"
)

type SignUpJSON struct {
	Email                string `json:"email" binding:"required"`
	Password             string `json:"password" binding:"required"`
	ConfirmationPassword string `json:"confirmationPassword" binding:"required"`
}

type SignInJSON struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json SignUpJSON
		if c.BindJSON(&json) != nil {
			c.JSON(400, lib.ErrorResponse("invalid Params"))
			return
		}
		email := json.Email
		password := json.Password
		confirmationPassword := json.ConfirmationPassword

		if password != confirmationPassword {
			c.JSON(400, lib.ErrorResponse("Password Mismatch"))
			return
		}
		encryptedPassword, err := lib.GetEncryptedPassword(password)
		if err != nil || encryptedPassword == "" {
			c.JSON(400, lib.ErrorResponse("encrypt error"))
			return
		}
		err = model.User{Email: email, Password: encryptedPassword}.Create()
		if err != nil {
			c.JSON(400, lib.ErrorResponse(err.Error()))
			return
		}
		c.AbortWithStatus(200)
	}
}

func SignIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json SignInJSON
		if c.BindJSON(&json) != nil {
			c.JSON(400, lib.ErrorResponse("Invalid Params"))
			return
		}
		email := json.Email
		password := json.Password
		user := model.NewUser(email).FindOne()
		if user == nil {
			c.JSON(400, lib.ErrorResponse("User not found"))
			return
		}
		if lib.ComparePassword(password, user.Password) {
			if tokenString, err := lib.GenerateTokenString(user.ID); err == nil {
				c.JSON(200, gin.H{"token": tokenString})
				return
			} else {
				c.JSON(400, lib.ErrorResponse("token generate error"))
				return
			}
		} else {
			c.JSON(400, lib.ErrorResponse("Invalid email or password"))
			return
		}
	}
}
