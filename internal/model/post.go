package model

import (
	"log"
	"time"
)

type Post struct {
	ID        int64     `xorm:"id pk" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    int64     `xorm:"user_id" json:"userId"`
	CreatedAt time.Time `xorm:"created" json:"createdAt"`
	UpdatedAt time.Time `xorm:"updated" json:"updatedAt"`
	Comment   Comment   `xorm:"extends"`
	Comments  []Comment
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
	var rPosts []Post
	engine.Table("post").Join("INNER", "comment", "post.id = comment.post_id").Find(&posts)
	for _, post := range posts {
		log.Println(post.Comments)
		has := false
		for i, result := range rPosts {
			if post.ID == result.ID {
				rPosts[i].Comments = append(rPosts[i].Comments, post.Comment)
				has = true
				break
			}
		}
		if !has {
			tmp := post
			tmp.Comments = append(tmp.Comments, tmp.Comment)
			rPosts = append(rPosts, tmp)
		}
	}
	return &rPosts
}

func (p Post) Exist() bool {
	has, _ := engine.Get(&p)
	return has
}
