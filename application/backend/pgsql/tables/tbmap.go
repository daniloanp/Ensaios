package tables

import (
	"github.com/daniloanp/Ensaios/application/backend/model/tables"
)

type  tableMap struct {
	userAccount                tables.UserAccount
	userPassword               tables.UserPassword
	userEmail                  tables.UserEmail
	userPersonalInformation    tables.UserPersonalInformation
	module                     tables.Module
	operation                  tables.Operation
	permission                 tables.Permission
	role                       tables.Role
}

func (d *tableMap) UserAccount() tables.UserAccount {
	return d.userAccount
}

func (d *tableMap) UserPassword() tables.UserPassword {
	return d.userPassword
}
func (d *tableMap) UserEmail() tables.UserEmail {
	return d.userEmail
}
func (d *tableMap) UserPersonalInformation() tables.UserPersonalInformation {
	return d.userPersonalInformation
}

func (d *tableMap) Module() tables.Module {
	return d.module
}

func (d *tableMap) Operation() tables.Operation {
	return d.operation
}

func (d *tableMap) Permission() tables.Permission {
	return d.permission
}

func (d *tableMap) Role() tables.Role {
	return d.role
}

