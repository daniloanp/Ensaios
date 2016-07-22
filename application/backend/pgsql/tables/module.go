package tables

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"database/sql"
	"github.com/daniloanp/Ensaios/application/backend/pgsql/conn"
	"github.com/daniloanp/Ensaios/application/backend/model/tables"
)

type module struct {
	create     *sql.Stmt
	getById    *sql.Stmt
	deleteById *sql.Stmt
	update     *sql.Stmt
}

func (m *module) Create(data *tables.ModuleData) (err error) {
	if data == nil {
		return model.ErrDataIsNil
	}
	if m.create == nil {
		const insQuery = `INSERT INTO "controller"."module"("name", "parent_module_id") VALUES($1, $2) returning "id"`
		m.create, err = conn.Db().Prepare(insQuery)
		if err != nil {
			return err
		}
	}
	row := m.create.QueryRow(data.Name, data.ParentModuleId)
	return row.Scan(&data.Id)
}

func (m *module) GetById(id int64) (data *tables.ModuleData, err error) {
	if m.getById == nil {
		const selQuery = `SELECT "id", "name", "parent_module_id" FROM "controller"."module" WHERE "id"=$1`
		m.getById, err = conn.Db().Prepare(selQuery)
		if err != nil {
			return nil, err
		}
	}
	row := m.getById.QueryRow(id)
	data = new(tables.ModuleData)
	err = row.Scan(
		&data.Id,
		&data.Name,
		&data.ParentModuleId,
	)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (m *module) DeleteById(id int64) (err error) {
	if m.deleteById == nil {
		const delQuery = `DELETE FROM "controller"."module" WHERE "id"=$1`
		m.deleteById, err = conn.Db().Prepare(delQuery)
		if err != nil {
			return err
		}
	}
	_, err = m.deleteById.Exec(id)
	return err
}

func (m *module) Update(data *tables.ModuleData) (err error) {
	if data == nil {
		return model.ErrDataIsNil
	}
	if m.update == nil {
		const uptQuery = `UPDATE "controller"."module" SET "name"=$1, "parent_module_id"=$2 WHERE "id"=$3`
		m.update, err = conn.Db().Prepare(uptQuery)
		if err != nil {
			return err
		}
	}
	_, err = m.update.Exec(data.Name, data.ParentModuleId, data.Id)
	return err
}

