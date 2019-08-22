package controller

import (
	"github.com/stobita/golang-sns-sample/internal/model"
)

type controller struct {
	repository repository
}

type repository interface {
	CreateUser(m *model.User) error
	UpdateUser(m *model.User) error
	FindUser(id uint) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
	GetUserPosts(userID uint) ([]*model.Post, error)

	CreatePost(m *model.Post) error
	GetPosts() ([]*model.Post, error)
	FindPost(id uint) (*model.Post, error)
	CreatePostComment(*model.Comment) error
}

func New(r repository) *controller {
	return &controller{
		repository: r,
	}
}
