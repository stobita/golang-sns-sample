package repository

import (
	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}
