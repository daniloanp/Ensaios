package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"fmt"
)

const defaultDbURL = "postgres://postgres:postgres@127.0.0.1/ensaios?sslmode=disable"

var pg (*sql.DB) = nil

func fatalOnError(err error) bool {
	if err != nil {
		log.Fatalln(err)
		return true
	}

	return false
}

func init() {
	var err error
	pg, err = sql.Open("postgres", defaultDbURL)
	fatalOnError(err)
}

func main() {
	tbs, err := loadTables()
	if fatalOnError(err) {
		return
	}
	fmt.Println(tbs)
}
