package pgsql

import "github.com/daniloanp/Ensaios/application/backend/model"

type  tableMap struct {
	userAccount                model.UserAccount
	userPassword               model.UserPassword
	userEmail                  model.UserEmail
	userPersonalInformation    model.UserPersonalInformation
	module                     model.Module
	operation                  model.Operation
	permission                 model.Permission
	role                       model.Role
	operationPermissionManager model.OperationPermissionManager
	permissionRoleManager      model.PermissionRoleManager
}

func (d *tableMap) UserAccount() model.UserAccount {
	return d.userAccount
}

func (d *tableMap) UserPassword() model.UserPassword {
	return d.userPassword
}
func (d *tableMap) UserEmail() model.UserEmail {
	return d.userEmail
}
func (d *tableMap) UserPersonalInformation() model.UserPersonalInformation {
	return d.userPersonalInformation
}

func (d *tableMap) Module() model.Module {
	return d.module
}

func (d *tableMap) Operation() model.Operation {
	return d.operation
}

func (d *tableMap) Permission() model.Permission {
	return d.permission

}

func (d *tableMap) Role() model.Role {
	return d.role
}

func (d *tableMap) OperationPermissionManager() model.OperationPermissionManager {
	return d.operationPermissionManager
}

func (d *tableMap) PermissionRoleManager() model.PermissionRoleManager {
	return d.permissionRoleManager
}

