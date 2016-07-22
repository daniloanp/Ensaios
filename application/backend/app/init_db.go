package app

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"github.com/daniloanp/Ensaios/application/backend/pgsql/tables"
)

var (
	db model.Tb = nil
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


func getDb() model.Tb {
	return tables.Instance()
}

func Db() model.Tb {
	if db == nil {
		db = getDb()
	}
	return db
}
