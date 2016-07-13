package tables

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"database/sql"
	"github.com/daniloanp/Ensaios/application/backend/pgsql/conn"
	"github.com/daniloanp/Ensaios/application/backend/model/tables"
)

type operation struct {
	create     *sql.Stmt
	getByID    *sql.Stmt
	deleteByID *sql.Stmt
	update     *sql.Stmt
}

func (op *operation) Create(data *tables.OperationData) (err error) {
	if data == nil {
		return model.ErrDataIsNil
	}
	if op.create == nil {
		const insQuery = `INSERT INTO "controller"."operation"("name", "module_id") VALUES($1, $2) returning "id"`
		op.create, err = conn.Db().Prepare(insQuery)
		if err != nil {
			return err
		}
	}

	row := op.create.QueryRow(data.Name, data.ModuleID)
	return row.Scan(&data.ID)
}

func (op *operation) GetByID(id int64) (data *tables.OperationData,err  error) {
	if op.getByID == nil {
		const selQuery = `SELECT "id", "name", "module_id" FROM "controller"."operation" WHERE "id"=$1`
		op.getByID, err = conn.Db().Prepare(selQuery)
		if err != nil {
			return nil, err
		}
	}
	row := op.getByID.QueryRow(id)
	data = new(tables.OperationData)
	err = row.Scan(
		&data.ID,
		&data.Name,
		&data.ModuleID,
	)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (op *operation) DeleteByID(id int64) (err error) {
	if op.deleteByID == nil {
		const delQuery = `DELETE FROM "controller"."operation" WHERE "id"=$1`
		op.deleteByID, err = conn.Db().Prepare(delQuery)
		if err != nil {
			return err
		}
	}
	_, err = op.deleteByID.Exec(id)
	return err
}

func (op *operation) Update(data *tables.OperationData) (err error) {
	if data == nil {
		return model.ErrDataIsNil
	}
	if op.update == nil {
		const uptQuery = `UPDATE "controller"."operation" SET "name"=$1, "module_id"=$2 WHERE "id"=$3`
		op.update, err = conn.Db().Prepare(uptQuery)
		if err != nil {
			return err
		}
	}
	_, err = op.update.Exec(data.Name, data.ModuleID, data.ID)
	return err
}
