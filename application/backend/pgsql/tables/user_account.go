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
	getByID    *sql.Stmt
	deleteByID *sql.Stmt
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
	return row.Scan(&data.ID)
}

func (ua *userAccount) GetByUsername(username string) (data *tables.UserAccountData, err error) {
	if ua.getByUsername== nil {
		const selQuery = `SELECT "id", "username", "registration_datetime" FROM "users"."user_account" WHERE "username"=$1`
		ua.getByUsername, err = conn.Db().Prepare(selQuery)
		if err != nil {
			return nil, err
		}
	}
	row := ua.getByID.QueryRow(username)
	data = new(tables.UserAccountData)
	err = row.Scan(
		&data.ID,
		&data.Username,
		&data.RegistrationDateTime,
	)

	if err != nil {
		return nil, err
	}
	return
}

func (ua *userAccount) GetByID(id int64) (data *tables.UserAccountData, err error) {
	if ua.getByID == nil {
		const selQuery = `SELECT "id", "username", "registration_datetime" FROM "users"."user_account" WHERE "id"=$1`
		ua.getByID, err = conn.Db().Prepare(selQuery)
		if err != nil {
			return nil, err
		}
	}
	row := ua.getByID.QueryRow(id)
	data = new(tables.UserAccountData)
	err = row.Scan(
		&data.ID,
		&data.Username,
		&data.RegistrationDateTime,
	)
	if err != nil {
		return nil, err
	}
	return
}

func (ua *userAccount) DeleteByID(id int64) (err error) {
	if ua.deleteByID == nil {
		const delQuery = `DELETE FROM "users"."user_account" WHERE "id"=$1`
		ua.deleteByID, err = conn.Db().Prepare(delQuery)
		if err != nil {
			return err
		}
	}
	_, err = ua.deleteByID.Exec(id)
	return err
}

func (_ *userAccount) Update(userAccountData *tables.UserAccountData) (err error) {
	return nil
}







