package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/stobita/golang-sns-sample/internal/model"
	"github.com/stobita/golang-sns-sample/internal/presenter"
)

type createPostRequestBody struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type createPostCommentRequestBody struct {
	Content string `json:"content" binding:"required"`
}

func (c *controller) CreatePost() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var json createPostRequestBody
		if err := ctx.ShouldBindJSON(&json); err != nil {
			log.Error(err)
			ctx.JSON(http.StatusBadRequest, presenter.ErrorResponse("Invalid params"))
			return
		}
		title := json.Title
		content := json.Content
		userID := ctx.Keys["userId"].(uint)
		model := &model.Post{
			Title:   title,
			Content: content,
			UserID:  userID,
		}
		if err := c.repository.CreatePost(model); err != nil {
			log.Error(err)
			ctx.JSON(http.StatusInternalServerError, presenter.ErrorResponse("Failed create post"))
			return
		}

		ctx.AbortWithStatus(http.StatusOK)
	}
}

func (c *controller) GetPosts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		posts, err := c.repository.GetPosts()
		if err != nil {
			log.Error(err)
			ctx.JSON(http.StatusInternalServerError, presenter.ErrorResponse("Failed get posts"))
			return
		}
		res, err := presenter.PostsResponse(posts)
		if err != nil {
			log.Error(err)
			ctx.JSON(http.StatusInternalServerError, presenter.ErrorResponse("response data error"))
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *controller) CreatePostComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var json createPostCommentRequestBody
		if err := ctx.ShouldBindJSON(&json); err != nil {
			log.Error(err)
			ctx.JSON(http.StatusBadRequest, presenter.ErrorResponse("Invalid params"))
			return
		}
		content := json.Content
		postID, _ := strconv.Atoi(ctx.Param("postId"))
		userID := ctx.Keys["userId"].(uint)

		post, err := c.repository.FindPost(uint(postID))
		if err != nil || post == nil {
			log.Error(err)
			ctx.JSON(http.StatusBadRequest, presenter.ErrorResponse("Post not found"))
			return
		}

		comment := &model.Comment{
			Content: content,
			PostID:  post.ID,
			UserID:  userID,
		}

		if err := c.repository.CreatePostComment(comment); err != nil {
			log.Error(err)
			ctx.JSON(http.StatusInternalServerError, presenter.ErrorResponse("Failed create comment"))
			return
		}

		ctx.AbortWithStatus(http.StatusOK)
	}
}
