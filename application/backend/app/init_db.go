package app

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"github.com/daniloanp/Ensaios/application/backend/model/pgsql"
)

var (
	db model.Db = nil
)

const (
	ErrorOnInitDb = "Error on InitDb"
	AnonymousRole = int64(-1)
	AdminRole = int64(0)
)

func panicOnError(err error) {
	if err != nil {
		panic(ErrorOnInitDb)
	}
}


func getDb() model.Db {
	return pgsql.Instance()
}

func Db() model.Db {
	if db == nil {
		db = getDb()
	}
	return db
}
