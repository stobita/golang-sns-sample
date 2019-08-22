package server

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stobita/golang-sns-sample/internal/controller"
	"github.com/stobita/golang-sns-sample/internal/db"
	"github.com/stobita/golang-sns-sample/internal/middleware"
	"github.com/stobita/golang-sns-sample/internal/repository"
)

var defaultPort = "8080"

func Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	db := db.NewGormConn()
	defer db.Close()
	engine, err := getEngine(db)
	if err != nil {
		return err
	}
	return engine.Run(":" + port)
}

func getEngine(db *gorm.DB) (*gin.Engine, error) {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	repository := repository.New(db)
	controller := controller.New(repository)

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
	return r, nil
}
