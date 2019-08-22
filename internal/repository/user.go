package repository

import "github.com/stobita/golang-sns-sample/internal/model"

func (r *repository) CreateUser(m *model.User) error {
	if err := r.db.Create(m).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateUser(m *model.User) error {
	if err := r.db.Update(m).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) FindUser(id uint) (*model.User, error) {
	user := &model.User{}
	if err := r.db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) FindUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) GetUserPosts(userID uint) ([]*model.Post, error) {
	var posts []*model.Post
	if err := r.db.Where("user_id = ?", userID).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
