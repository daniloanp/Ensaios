package tables

import (
	"github.com/daniloanp/Ensaios/application/backend/model/tables"
)

type userPersonalInformation struct{}

func (_ *userPersonalInformation) Create(userPersonalInformation *tables.UserPersonalInformation) (err error) {
	return nil
}
func (_ *userPersonalInformation) Update(userPersonalInformation  *tables.UserPersonalInformation) (err error) {
	return nil
}
func (_ *userPersonalInformation) DeleteById(id int64) (err error) {
	return nil
}

func (_ *userPersonalInformation) GetById(id int64) (*tables.UserPersonalInformation, error) {
	return nil, nil
}

