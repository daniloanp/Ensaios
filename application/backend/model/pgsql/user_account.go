package pgsql

import "github.com/daniloanp/Ensaios/application/backend/model"

type userAccount struct{}

func (_ *userAccount) Create(userAccountData *model.UserAccountData) (err error) {
	return nil
}
func (_ *userAccount) GetByUsername(username string) (*model.UserAccountData, error) {
	return nil,nil
}
func (_ *userAccount) GetByID(id int64) (*model.UserAccountData, error) {
	return nil,nil
}
func (_ *userAccount) DeleteByID(id int64) (err error) {
	return nil
}
func (_ *userAccount) Update(userAccountData *model.UserAccountData) (err error) {
	return nil
}





