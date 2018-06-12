package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stobita/golang-sns-sample/handler"
	"github.com/stobita/golang-sns-sample/middleware"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.GET("/", handler.Root())
	r.POST("signup", handler.SignUp())
	r.POST("signin", handler.SignIn())
	r.GET("/posts", handler.GetPosts())
	authorized := r.Group("/", middleware.TokenAuthMiddleware())
	{
		authorized.PUT("/user", handler.UserUpdate())
		authorized.POST("/post", handler.CreatePost())
		authorized.GET("/user/posts", handler.GetUserPosts())
		authorized.POST("/post/:postId/comment", handler.CreatePostComment())
	}
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
