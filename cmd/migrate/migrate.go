package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/stobita/golang-sns-sample/internal/db"
	"github.com/stobita/golang-sns-sample/internal/model"
)

func main() {
	db := db.NewGormConn()
	db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Comment{},
	)
}
