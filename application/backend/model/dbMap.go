package model


//TbMap is a set of db mapping. TODO:It should be a interface
type dbMap struct {
	UserAccount                UserAccount
	UserPassword               UserPassword
	UserEmail                  UserEmail
	UserPersonalInformation    UserPersonalInformation
	Module                     Module
	Operation                  Operation
	Permission                 Permission
	Role                       Role
	OperationPermissionManager OperationPermissionManager
	PermissionRoleManager      PermissionRoleManager
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


