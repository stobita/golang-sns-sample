package main

import (
	"database/sql"
	"flag"
	"fmt"
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
	dbUser := os.Getenv("SAMPLE_DB_USER")
	dbPassword := os.Getenv("SAMPLE_DB_PASS")
	dbHost := os.Getenv("SAMPLE_DB_HOST")
	dbName := os.Getenv("SAMPLE_DB_NAME")
	db, err := sql.Open(driver, fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName))
	if err != nil {
		log.Fatalf("goose run: %v", err)
	}
	if err := goose.Run(command, db, dir); err != nil {
		log.Fatalf("goose run: %v", err)
	}
}
