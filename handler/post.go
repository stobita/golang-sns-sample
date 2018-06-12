package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stobita/golang-sns-sample/lib"
	"github.com/stobita/golang-sns-sample/model"
)

type PostJSON struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type CommentJSON struct {
	Content string `json:"content" binding:"required"`
}

func CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json PostJSON
		if c.BindJSON(&json) != nil {
			c.JSON(400, lib.ErrorResponse("Invalid Params"))
			return
		}
		title := json.Title
		content := json.Content
		userID := c.Keys["userId"].(int64)
		err := model.NewPost(title, content, userID).Create()
		if err != nil {
			c.JSON(400, lib.ErrorResponse(err.Error()))
			return
		}
		c.AbortWithStatus(200)
	}
}

func GetPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := model.Post{}.GetAll()
		c.JSON(200, result)
	}
}

func CreatePostComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json CommentJSON
		if c.Bind(&json) != nil {
			c.JSON(400, lib.ErrorResponse("Invalid Params"))
			return
		}
		content := json.Content
		postID, _ := strconv.ParseInt(c.Param("postId"), 10, 64)
		userID := c.Keys["userId"].(int64)
		if (!model.Post{ID: postID}.Exist()) {
			c.JSON(400, lib.ErrorResponse("Invalid PostId"))
			return
		}
		err := model.Comment{Content: content, PostID: postID, UserID: userID}.Create()
		if err != nil {
			c.JSON(400, lib.ErrorResponse(err.Error()))
			return
		}
		c.AbortWithStatus(200)
	}
}
