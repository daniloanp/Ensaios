package pgsql

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"database/sql"
	"github.com/sosedoff/pgweb/pkg/data"
)

type permission struct{
	create     *sql.Stmt
	getByID    *sql.Stmt
	deleteByID *sql.Stmt
	update     *sql.Stmt
}

func (perm *permission) Create(data *model.PermissionData) (err error) {
	if data == nil {
		return model.ErrDataIsNil
	}
	if perm.create == nil {
		const insQuery = `INSERT INTO "controller"."permission"("description") VALUES($1) returning "id"`
		perm.create, err = pgSql.Prepare(insQuery)
		if err != nil {
			return err
		}
	}

	row := perm.create.QueryRow(data.Description)
	return row.Scan(&data.ID)
}
func (perm *permission) GetByID(id int64) (data *model.PermissionData, err error) {
	if perm.getByID == nil {
		const selQuery = `SELECT "id", "description" FROM "controller"."permission" WHERE "id"=$1`
		perm.getByID, err = pgSql.Prepare(selQuery)
		if err != nil {
			return nil, err
		}
	}
	row := perm.getByID.QueryRow(id)
	data = new(model.ModuleData)
	err = row.Scan(
		&data.ID,
		&data.Description,
	)
	if err != nil {
		return nil, err
	}
	return data, err
}
func (perm *permission) DeleteByID(id int64) (err error) {
	if perm.deleteByID == nil {
		const delQuery = `DELETE FROM "controller"."permission" WHERE "id"=$1`
		perm.deleteByID, err = pgSql.Prepare(delQuery)
		if err != nil {
			return err
		}
	}
	_, err = perm.deleteByID.Exec(id)
	return err
}
func (perm *permission) Update(data *model.PermissionData) (err error) {
	if perm.update == nil {
		const uptQuery = `UPDATE "controller"."permission" SET "description"=$1 WHERE "id"=$3`
		perm.update, err = pgSql.Prepare(uptQuery)
		if err != nil {
			return err
		}
	}
	_, err = perm.update.Exec(data.Description, data.ID)
	return err
}