package pgsql

import (
	"database/sql"
	"log"
	//Necessary because pq does implement things expected by "database/sql" package
	_ "github.com/lib/pq"
)

const dbURL = "postgres://postgres:a9*@0.0.0.0/ensaios?sslmode=disable"

var pgSql (*sql.DB)

func init() {
	var err error
	pgSql, err = sql.Open("postgres", dbURL)
	pgSql.SetMaxOpenConns(3)
	pgSql.SetMaxIdleConns(3)
	if err != nil {
		log.Fatal(err)
	}
}
