package model

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine
var err error

func init() {
	driver := "mysql"
	dbUser := os.Getenv("SAMPLE_DB_USER")
	dbPassword := os.Getenv("SAMPLE_DB_PASS")
	dbHost := os.Getenv("SAMPLE_DB_POST")
	dbName := os.Getenv("SAMPLE_DB_NAME")
	engine, err = xorm.NewEngine(driver, fmt.Sprintf("%d:%d@tcp(%d)/%d", dbUser, dbPassword, dbHost, dbName))
	if err != nil {
		panic("failed to connect database")
	}
	engine.ShowSQL(true)
	defer engine.Clone()
}
