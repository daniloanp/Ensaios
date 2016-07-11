package pgsql

import "github.com/daniloanp/Ensaios/application/backend/model"

type role struct{}

func (_ *role) Create(roleData *model.RoleData) (err error) {
	return nil
}
func (_ *role) GetByID(id int64) (*model.RoleData, error) {
	return nil,nil
}
func (_ *role) DeleteByID(id int64) (err error) {
	return nil
}
func (_ *role) Update(roleData *model.RoleData) (err error) {
	return nil
}

func (_ *role) HasPermission(roleData *model.RoleData, operationPath string) (bool, error) {
	return false, nil
}

