package tables

import (
	"github.com/daniloanp/Ensaios/application/backend/model/tables"
)

type userPassword struct{}

func (_ *userPassword) Create(userPasswordData *tables.UserPasswordData) (err error) {
	return nil
}
func (_ *userPassword) Update(userPasswordData *tables.UserPasswordData) (err error) {
	return nil
}
func (_ *userPassword) DeleteByID(id int64) (err error) {
	return nil
}
func (_ *userPassword) GetByID(id int64) (*tables.UserPasswordData, error) {
	return nil,nil
}
