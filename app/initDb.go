package app

import (
	"github.com/daniloanp/Ensaios/app/backend/model"
)

var (
	db model.DbMap
)

const (
	ErrorOnInitDb = "Error on InitDb"
)

func initDb() {
	var err error
	var baseModule = &model.ModuleData{
		Name:"",
	}
	err = db.Module.Create(baseModule)
	if err != nil {
		panic(ErrorOnInitDb)
	}




}
