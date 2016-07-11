package model



type Db interface {
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


