package app

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"github.com/daniloanp/Ensaios/application/backend/modelPgsql"
)

var (
	db model.DbMap = nil
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


func getDb() model.DbMap {
	return modelPgsql.Instance()
}

func Db() model.DbMap {
	if db == nil {
		db = getDb()
	}
	return db
}
