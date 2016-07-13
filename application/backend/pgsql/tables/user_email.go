package tables

import (
	"github.com/daniloanp/Ensaios/application/backend/model/tables"
)

type userEmail struct{}

func (_ *userEmail) Create(userEmailData *tables.UserEmailData) (err error) {
	return nil
}

func (_ *userEmail) Update(userEmailData  *tables.UserEmailData) (err error) {
	return nil
}

func (_ *userEmail) DeleteByAddress(address string) (err error) {
	return nil
}

func (_ *userEmail) GetByAddress(address string) (*tables.UserEmailData, error) {
	return nil,nil
}
