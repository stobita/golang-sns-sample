package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine
var err error

func init() {
	engine, err = xorm.NewEngine("mysql", "gouser:gouser@/golang_sns_sample")
	if err != nil {
		panic("failed to connect database")
	}
	engine.ShowSQL(true)
	defer engine.Clone()
}
