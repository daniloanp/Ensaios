package bootstrap

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"database/sql"
)

type (
	operation struct {
		model.OperationData
		Module *module
	}
	module struct {
		model.ModuleData
		Operations   []*operation
		ParentModule *module
	}
)

func (op *operation) insert() error {
	err := db.Operation().Create(&op.OperationData)
	return err
}

func (mod *module) insert() error {
	if mod.ParentModule != nil {
		mod.ParentModuleID = sql.NullInt64{
			Valid: true,
			Int64:mod.ParentModule.ID,
		}
	}
	if err := db.Module().Create(&mod.ModuleData); err != nil {
		return err
	}

	if mod.Operations != nil {
		for _, op := range mod.Operations {
			op.Module = mod
			op.ModuleID = mod.ID
			if err := op.insert(); err != nil {
				return err
			}
		}
	}
	return nil
}