package pgsql

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"database/sql"
)

type operation struct {
	create     *sql.Stmt
	getByID    *sql.Stmt
	deleteByID *sql.Stmt
	update     *sql.Stmt
}



func (op *operation) Create(data *model.OperationData) (err error) {
	if data == nil {
		return model.ErrDataIsNil
	}
	if op.create == nil {
		const insQuery = `INSERT INTO controller.module(name, parent_module_id) VALUES($1, $2) returning id`
		op.create, err = pgSql.Prepare(insQuery)
		if err != nil {
			return err
		}
	}


	return nil
}
func (op *operation) GetByID(id int64) (*model.OperationData, error) {
	return nil, nil
}
func (op *operation) DeleteByID(id int64) (err error) {
	return nil
}
func (op *operation) Update(data *model.OperationData) (err error) {
	if data == nil {
		return model.ErrDataIsNil
	}
	return nil
}
