package model


//TbMap is a set of db mapping. TODO:It should be a interface
type dbMap struct {
	userAccount                UserAccount
	userPassword               UserPassword
	userEmail                  UserEmail
	userPersonalInformation    UserPersonalInformation
	module                     Module
	operation                  Operation
	permission                 Permission
	role                       Role
	operationPermissionManager OperationPermissionManager
	permissionRoleManager      PermissionRoleManager
}

type DbMap interface {
	UserAccount() UserAccount
	UserPassword() UserPassword
	UserEmail() UserEmail
	UserPersonalInformation() UserPersonalInformation
	Module() Module
	Operation() Operation
	Permission() Permission
	Role() Role
	OperationPermissionManager() OperationPermissionManager
	PermissionRoleManager() PermissionRoleManager
}


