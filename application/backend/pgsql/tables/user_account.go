package tables

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"database/sql"
	"github.com/daniloanp/Ensaios/application/backend/pgsql/conn"
	"github.com/daniloanp/Ensaios/application/backend/model/tables"
)

type userAccount struct{
	create     *sql.Stmt
	getByUsername    *sql.Stmt
	getById    *sql.Stmt
	deleteById *sql.Stmt
	update     *sql.Stmt
}

func (ua *userAccount) Create(data *tables.UserAccountData) (err error) {
	if data == nil {
		return model.ErrDataIsNil
	}
	if ua.create == nil {
		const insQuery = `INSERT INTO "users"."user_account" ("username") VALUES($1) returning "id"`
		ua.create, err = conn.Db().Prepare(insQuery)
		if err != nil {
			return err
		}
	}

	row := ua.create.QueryRow(data.Username)
	return row.Scan(&data.Id)
}

func (ua *userAccount) GetByUsername(username string) (data *tables.UserAccountData, err error) {
	if ua.getByUsername== nil {
		const selQuery = `SELECT "id", "username", "registration_datetime" FROM "users"."user_account" WHERE "username"=$1`
		ua.getByUsername, err = conn.Db().Prepare(selQuery)
		if err != nil {
			return nil, err
		}
	}
	row := ua.getById.QueryRow(username)
	data = new(tables.UserAccountData)
	err = row.Scan(
		&data.Id,
		&data.Username,
		&data.RegistrationDateTime,
	)

	if err != nil {
		return nil, err
	}
	return
}

func (ua *userAccount) GetById(id int64) (data *tables.UserAccountData, err error) {
	if ua.getById == nil {
		const selQuery = `SELECT "id", "username", "registration_datetime" FROM "users"."user_account" WHERE "id"=$1`
		ua.getById, err = conn.Db().Prepare(selQuery)
		if err != nil {
			return nil, err
		}
	}
	row := ua.getById.QueryRow(id)
	data = new(tables.UserAccountData)
	err = row.Scan(
		&data.Id,
		&data.Username,
		&data.RegistrationDateTime,
	)
	if err != nil {
		return nil, err
	}
	return
}

func (ua *userAccount) DeleteById(id int64) (err error) {
	if ua.deleteById == nil {
		const delQuery = `DELETE FROM "users"."user_account" WHERE "id"=$1`
		ua.deleteById, err = conn.Db().Prepare(delQuery)
		if err != nil {
			return err
		}
	}
	_, err = ua.deleteById.Exec(id)
	return err
}

func (_ *userAccount) Update(userAccountData *tables.UserAccountData) (err error) {
	return nil
}







