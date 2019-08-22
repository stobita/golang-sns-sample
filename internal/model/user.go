package model

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
