package model


//TbMap is a set of db mapping.
type DbMap struct {
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


