package tables

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"database/sql"
	"github.com/daniloanp/Ensaios/application/backend/pgsql/conn"
	"github.com/daniloanp/Ensaios/application/backend/model/tables"
)

type permission struct{
	create     *sql.Stmt
	getByID    *sql.Stmt
	deleteByID *sql.Stmt
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
	return row.Scan(&data.ID)
}

func (per *permission) GetByID(id int64) (data *tables.PermissionData, err error) {
	if per.getByID == nil {
		const selQuery = `SELECT "id", "description" FROM "controller"."permission" WHERE "id"=$1`
		per.getByID, err = conn.Db().Prepare(selQuery)
		if err != nil {
			return nil, err
		}
	}
	row := per.getByID.QueryRow(id)
	data = new(tables.PermissionData)
	err = row.Scan(
		&data.ID,
		&data.Description,
	)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (per *permission) DeleteByID(id int64) (err error) {
	if per.deleteByID == nil {
		const delQuery = `DELETE FROM "controller"."permission" WHERE "id"=$1`
		per.deleteByID, err = conn.Db().Prepare(delQuery)
		if err != nil {
			return err
		}
	}
	_, err = per.deleteByID.Exec(id)
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
	_, err = per.update.Exec(data.Description, data.ID)
	return err
}