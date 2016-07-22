package tables

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"database/sql"
	"github.com/daniloanp/Ensaios/application/backend/pgsql/conn"
	"github.com/daniloanp/Ensaios/application/backend/model/tables"
)

type permission struct{
	create     *sql.Stmt
	getById    *sql.Stmt
	deleteById *sql.Stmt
	update     *sql.Stmt
}

func (per *permission) Create(data *tables.PermissionData) (err error) {
	if data == nil {
		return model.ErrDataIsNil
	}
	if per.create == nil {
		const insQuery = `INSERT INTO "controller"."permission"("description") VALUES($1) returning "id"`
		per.create, err = conn.Db().Prepare(insQuery)
		if err != nil {
			return err
		}
	}

	row := per.create.QueryRow(data.Description)
	return row.Scan(&data.Id)
}

func (per *permission) GetById(id int64) (data *tables.PermissionData, err error) {
	if per.getById == nil {
		const selQuery = `SELECT "id", "description" FROM "controller"."permission" WHERE "id"=$1`
		per.getById, err = conn.Db().Prepare(selQuery)
		if err != nil {
			return nil, err
		}
	}
	row := per.getById.QueryRow(id)
	data = new(tables.PermissionData)
	err = row.Scan(
		&data.Id,
		&data.Description,
	)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (per *permission) DeleteById(id int64) (err error) {
	if per.deleteById == nil {
		const delQuery = `DELETE FROM "controller"."permission" WHERE "id"=$1`
		per.deleteById, err = conn.Db().Prepare(delQuery)
		if err != nil {
			return err
		}
	}
	_, err = per.deleteById.Exec(id)
	return err
}

func (per *permission) Update(data *tables.PermissionData) (err error) {
	if data == nil {
		return model.ErrDataIsNil
	}
	if per.update == nil {
		const uptQuery = `UPDATE "controller"."permission" SET "description"=$1 WHERE "id"=$3`
		per.update, err = conn.Db().Prepare(uptQuery)
		if err != nil {
			return err
		}
	}
	_, err = per.update.Exec(data.Description, data.Id)
	return err
}