package pgsql

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"database/sql"
)

type module struct {
	create *sql.Stmt
	getByID        *sql.Stmt
	deleteByID     *sql.Stmt
	update         *sql.Stmt
}

func (m *module) Create(data *model.ModuleData) (err error) {
	if data == nil {
		return model.ErrDataIsNil
	}
	if m.create == nil {
		const insQuery = `INSERT INTO "controller"."module"("name", "parent_module_id") VALUES($1, $2) returning "id"`
		m.create, err = pgSql.Prepare(insQuery)
		if err != nil {
			return err
		}
	}
	row := m.create.QueryRow(data.Name, data.ParentModuleID)
	return row.Scan(&data.ID)
}

func (m *module) GetByID(id int64) (data *model.ModuleData, err error) {
	if m.getByID == nil {
		const selQuery = `SELECT "id", "name", "parent_module_id" FROM "controller"."module" WHERE "id"=$1`
		m.getByID, err = pgSql.Prepare(selQuery)
		if err != nil {
			return nil, err
		}
	}
	row := m.getByID.QueryRow(id)
	data = new(model.ModuleData)
	err = row.Scan(
		&data.ID,
		&data.Name,
		&data.ParentModuleID,
	)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (m *module) DeleteByID(id int64) (err error) {
	if m.deleteByID == nil {
		const delQuery = `DELETE FROM "controller"."module" WHERE "id"=$1`
		m.deleteByID, err = pgSql.Prepare(delQuery)
		if err != nil {
			return err
		}
	}
	_, err = m.deleteByID.Exec(id)
	return err
}

func (m *module) Update(data *model.ModuleData) (err error) {
	if m.update == nil {
		const uptQuery = `UPDATE "controller"."module" SET "name"=$1, "parent_module_id"=$2 WHERE "id"=$3`
		m.update, err = pgSql.Prepare(uptQuery)
		if err != nil {
			return err
		}
	}
	_, err = m.update.Exec(data.Name, data.ParentModuleID, data.ID)
	return err
}
