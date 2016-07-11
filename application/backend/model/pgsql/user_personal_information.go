package pgsql

import "github.com/daniloanp/Ensaios/application/backend/model"

type userPersonalInformation struct{}

func (_ *userPersonalInformation) Create(userPersonalInformation *model.UserPersonalInformation) (err error) {
	return nil
}
func (_ *userPersonalInformation) Update(userPersonalInformation  *model.UserPersonalInformation) (err error) {
	return nil
}
func (_ *userPersonalInformation) DeleteById(id int64) (err error) {
	return nil
}

func (_ *userPersonalInformation) GetById(id int64) (*model.UserPersonalInformation, error) {
	return nil, nil
}

