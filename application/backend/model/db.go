package model

import "github.com/daniloanp/Ensaios/application/backend/model/tables"

type Db interface {
	UserAccount() tables.UserAccount
	UserPassword() tables.UserPassword
	UserEmail() tables.UserEmail
	UserPersonalInformation() tables.UserPersonalInformation
	Module() tables.Module
	Operation() tables.Operation
	Permission() tables.Permission
	Role() tables.Role
	//OperationPermissionManager() tables.OperationPermissionManager
	//PermissionRoleManager() tables.PermissionRoleManager
}


