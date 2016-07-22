package tables

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"database/sql"
	"github.com/daniloanp/Ensaios/application/backend/pgsql/conn"
	"github.com/daniloanp/Ensaios/application/backend/model/tables"
)

type role struct{
	create     *sql.Stmt
	getById    *sql.Stmt
	deleteById *sql.Stmt
	update     *sql.Stmt
}

func (ro *role) Create(data *tables.RoleData) (err error) {
	if data == nil {
		return model.ErrDataIsNil
	}
	if ro.create == nil {
		const insQuery = `INSERT INTO "controller"."role"("description", "parent_role_id") VALUES($1, $2) returning "id"`
		ro.create, err = conn.Db().Prepare(insQuery)
		if err != nil {
			return err
		}
	}

	row := ro.create.QueryRow(data.Description, data.ParentRoleId)
	return row.Scan(&data.Id)
}
func (ro *role) GetById(id int64) (data *tables.RoleData, err error) {
	if ro.getById == nil {
		const selQuery = `SELECT "id", "description", "parent_role_id" FROM "controller"."role" WHERE "id"=$1`
		ro.getById, err = conn.Db().Prepare(selQuery)
		if err != nil {
			return nil, err
		}
	}
	row := ro.getById.QueryRow(id)
	data = new(tables.RoleData)
	err = row.Scan(
		&data.Id,
		&data.Description,
		&data.ParentRoleId,
	)
	if err != nil {
		return nil, err
	}

	return data, err
}
func (ro *role) DeleteById(id int64) (err error) {
	if ro.deleteById == nil {
		const delQuery = `DELETE FROM "controller"."role" WHERE "id"=$1`
		ro.deleteById, err = conn.Db().Prepare(delQuery)
		if err != nil {
			return err
		}
	}
	_, err = ro.deleteById.Exec(id)
	return err
}

func (ro *role) Update(data *tables.RoleData) (err error) {
	if data == nil {
		return model.ErrDataIsNil
	}
	if ro.update == nil {
		const uptQuery = `UPDATE "controller"."role" SET "description"=$1, "parent_role_id"=$2 WHERE "id"=$3`
		ro.update, err = conn.Db().Prepare(uptQuery)
		if err != nil {
			return err
		}
	}
	_, err = ro.update.Exec(data.Description, data.ParentRoleId, data.Id)
	return err
}

func (_ *role) HasPermission(roleData *tables.RoleData, operationPath string) (bool, error) {
	return false, nil
}

