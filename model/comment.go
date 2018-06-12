package model

import "time"

type Comment struct {
	ID        int64     `xorm:"id pk" json:"id"`
	Content   string    `json:"content"`
	PostID    int64     `xorm:"post_id" json:"postId"`
	UserID    int64     `xorm:"user_id" json:"userId"`
	CreatedAt time.Time `xorm:"created" json:"createdAt"`
	UpdatedAt time.Time `xorm:"updated" json:"updatedAt"`
}

func (c Comment) Create() error {
	_, err := engine.Insert(&c)
	if err != nil {
		return err
	}
	return nil
}
