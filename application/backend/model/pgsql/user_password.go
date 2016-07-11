package pgsql

import "github.com/daniloanp/Ensaios/application/backend/model"

type userPassword struct{}

func (_ *userPassword) Create(userPasswordData *model.UserPasswordData) (err error) {
	return nil
}
func (_ *userPassword) Update(userPasswordData *model.UserPasswordData) (err error) {
	return nil
}
func (_ *userPassword) DeleteByID(id int64) (err error) {
	return nil
}
func (_ *userPassword) GetByID(id int64) (*model.UserPasswordData, error) {
	return nil,nil
}
