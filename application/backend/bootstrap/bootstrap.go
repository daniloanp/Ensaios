package bootstrap

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"github.com/daniloanp/Ensaios/application/backend/app"
)

var (
	db model.Tb = app.Db()
)

func Bootstrap() error {
	if err := baseModule.insert(); err != nil {
		return err
	}
	if err:= adminModule.insert(); err != nil {
		return err
	}
	for _, mod := range adminSubModules {
		if err:= mod.insert(); err != nil {
			return err
		}
	}
	return nil
}
