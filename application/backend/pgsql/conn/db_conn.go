package conn

import (
	"database/sql"
	"log"
	//Necessary because pq does implement things expected by "database/sql" package
	_ "github.com/lib/pq"
)

const defaultDbURL = "postgres://postgres:postgres@127.0.0.1/ensaios?sslmode=disable"

var postgres (*sql.DB) = nil

func Db() *sql.DB {
	if postgres == nil {
		var err error
		postgres, err = sql.Open("postgres", defaultDbURL)
		postgres.SetMaxOpenConns(6)
		postgres.SetMaxIdleConns(3)
		if err != nil {
			log.Println(err)
		}
	}
	return postgres
}
