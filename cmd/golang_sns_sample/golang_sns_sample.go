package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stobita/golang-sns-sample/internal/controller"
	"github.com/stobita/golang-sns-sample/internal/db"
	"github.com/stobita/golang-sns-sample/internal/middleware"
	"github.com/stobita/golang-sns-sample/internal/repository"
)

var defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db := db.NewGormConn()
	defer db.Close()
	repository := repository.New(db)
	controller := controller.New(repository)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	v1 := r.Group("/api/v1")
	{
		v1.POST("signup", controller.SignUp())
		v1.POST("signin", controller.SignIn())
		v1.GET("/posts", controller.GetPosts())
		authorized := v1.Group("/", middleware.TokenAuthMiddleware())
		{
			authorized.PUT("/user", controller.UpdateUser())
			authorized.POST("/post", controller.CreatePost())
			authorized.GET("/user/posts", controller.GetUserPosts())
			authorized.POST("/post/:postId/comment", controller.CreatePostComment())
		}
	}
	r.Run(":" + port)
}
