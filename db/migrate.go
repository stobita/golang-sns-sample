package main

import (
	"database/sql"
	"flag"
	"log"
	"os"

	"github.com/pressly/goose"

	_ "github.com/go-sql-driver/mysql"
)

var (
	flags  = flag.NewFlagSet("goose", flag.ExitOnError)
	dir    = "./db/migrations"
	driver = "mysql"
)

func main() {
	print(dir)
	flags.Parse(os.Args[1:])
	args := flags.Args()
	if len(args) < 1 {
		log.Fatalf("goose run: %v", "input goose action")
	}
	command := args[0]
	if command == "create" {
		if err := goose.Run("create", nil, dir, args[1], "sql"); err != nil {
			log.Fatalf("goose run: %v", err)
		}
		return
	}
	if len(args) > 1 {
		flags.Usage()
		return
	}
	if err := goose.SetDialect(driver); err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open(driver, "gouser:gouser@/golang_sns_sample")
	if err != nil {
		log.Fatalf("goose run: %v", err)
	}
	if err := goose.Run(command, db, dir); err != nil {
		log.Fatalf("goose run: %v", err)
	}
}
