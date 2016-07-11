package pgsql

import "github.com/daniloanp/Ensaios/application/backend/model"

type userEmail struct{}

func (_ *userEmail) Create(userEmailData *model.UserEmailData) (err error) {
	return nil
}
func (_ *userEmail) Update(userEmailData  *model.UserEmailData) (err error) {
	return nil
}
func (_ *userEmail) DeleteByAddress(address string) (err error) {
	return nil
}
func (_ *userEmail) GetByAddress(address string) (*model.UserEmailData, error) {
	return nil,nil
}
