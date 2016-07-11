package pgsql

import "github.com/daniloanp/Ensaios/application/backend/model"

type permission struct{}

func (_ *permission) Create(permissionData *model.PermissionData) (err error) {
	return nil
}
func (_ *permission) GetByID(id int64) (*model.PermissionData, error) {
	return nil, nil
}
func (_ *permission) DeleteByID(id int64) (err error) {
	return nil
}
func (_ *permission) Update(permissionData *model.PermissionData) (err error) {
	return nil
}