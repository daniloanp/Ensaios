package application

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
)

var (
	db model.DbMap = nil
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
