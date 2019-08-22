package model

import "time"

type Comment struct {
	ID        uint `gorm:"primary_key"`
	Content   string
	PostID    uint
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
