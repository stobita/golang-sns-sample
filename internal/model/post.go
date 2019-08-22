package model

import (
	"time"
)

type Post struct {
	ID        uint `gorm:"primary_key"`
	Title     string
	Content   string
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Comment   Comment
	Comments  []Comment
}
