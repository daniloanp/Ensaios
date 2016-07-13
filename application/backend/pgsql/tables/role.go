package tables

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"database/sql"
	"github.com/daniloanp/Ensaios/application/backend/pgsql/conn"
	"github.com/daniloanp/Ensaios/application/backend/model/tables"
)

type role struct{
	create     *sql.Stmt
	getByID    *sql.Stmt
	deleteByID *sql.Stmt
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

	row := ro.create.QueryRow(data.Description, data.ParentRoleID)
	return row.Scan(&data.ID)
}
func (ro *role) GetByID(id int64) (data *tables.RoleData, err error) {
	if ro.getByID == nil {
		const selQuery = `SELECT "id", "description", "parent_role_id" FROM "controller"."role" WHERE "id"=$1`
		ro.getByID, err = conn.Db().Prepare(selQuery)
		if err != nil {
			return nil, err
		}
	}
	row := ro.getByID.QueryRow(id)
	data = new(tables.RoleData)
	err = row.Scan(
		&data.ID,
		&data.Description,
		&data.ParentRoleID,
	)
	if err != nil {
		return nil, err
	}

	return data, err
}
func (ro *role) DeleteByID(id int64) (err error) {
	if ro.deleteByID == nil {
		const delQuery = `DELETE FROM "controller"."role" WHERE "id"=$1`
		ro.deleteByID, err = conn.Db().Prepare(delQuery)
		if err != nil {
			return err
		}
	}
	_, err = ro.deleteByID.Exec(id)
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
	_, err = ro.update.Exec(data.Description, data.ParentRoleID, data.ID)
	return err
}

func (_ *role) HasPermission(roleData *tables.RoleData, operationPath string) (bool, error) {
	return false, nil
}

