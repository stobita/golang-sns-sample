package model

import (
	"time"
)

// column name `ID` will be the default primary field

type User struct {
	ID        int64 `xorm:"id pk"`
	Email     string
	Password  string
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}

func (u User) Create() error {
	// Insert will return affected count
	_, err := engine.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}

func (u User) Update(userID int64) error {
	_, err := engine.Id(userID).Cols("email").Update(&u)
	if err != nil {
		return err
	}
	return nil
}

func (u User) FindOne() *User {
	has, _ := engine.Get(&u)
	if has {
		return &u
	}
	return nil
}

func (u User) GetPosts() *[]Post {
	var posts []Post
	engine.Table("post").
		Join("INNER", "user", "user.id = post.user_id").
		Where("user.id=?", u.ID).
		Find(&posts)
	return &posts
}
