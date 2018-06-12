package model

import "time"

type Post struct {
	ID        int64     `xorm:"id pk" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    int64     `xorm:"user_id" json:"userId"`
	CreatedAt time.Time `xorm:"created" json:"createdAt"`
	UpdatedAt time.Time `xorm:"updated" json:"updatedAt"`
}

func NewPost(title string, content string, userID int64) *Post {
	post := new(Post)
	post.Title = title
	post.Content = content
	post.UserID = userID
	return post
}

func (p Post) Create() error {
	_, err := engine.Insert(&p)
	if err != nil {
		return err
	}
	return nil
}

func (p Post) GetAll() *[]Post {
	var posts []Post
	engine.Find(&posts)
	return &posts
}

func (p Post) Exist() bool {
	has, _ := engine.Get(&p)
	return has
}
