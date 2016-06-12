package app

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
)

var (
	db *model.DbMap = nil
)

const (
	ErrorOnInitDb = "Error on InitDb"
	AnonymousRole = int64(-1)
	AdminRole = int64(0)
)

func throwPanic(err error) {
	if err != nil {
		panic(ErrorOnInitDb)
	}
}


func getDb() *model.DbMap {
	return nil // TODO:missing databse implementation
}

func Db() *model.DbMap {
	if db == nil {
		db = getDb()
	}
	return db
}
