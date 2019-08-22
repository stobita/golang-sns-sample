package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewGormConn() *gorm.DB {
	dbDriver := "mysql"
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASS")
	dbHost := os.Getenv("MYSQL_HOST")
	dbName := os.Getenv("MYSQL_DBNAME")
	conn, err := gorm.Open(dbDriver, fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=True", dbUser, dbPass, dbHost, dbName))
	if err != nil {
		log.Fatalln(err)
	}
	conn.SingularTable(true)
	return conn
}
