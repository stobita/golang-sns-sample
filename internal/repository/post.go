package repository

import "github.com/stobita/golang-sns-sample/internal/model"

func (r *repository) CreatePost(m *model.Post) error {
	if err := r.db.Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) GetPosts() ([]*model.Post, error) {
	var posts []*model.Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *repository) FindPost(id uint) (*model.Post, error) {
	post := &model.Post{}
	if err := r.db.First(post, id).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *repository) CreatePostComment(m *model.Comment) error {
	if err := r.db.Create(m).Error; err != nil {
		return err
	}
	return nil
}
