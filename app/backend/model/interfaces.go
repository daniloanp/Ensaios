package model

import (
	"database/sql"
)
type UserAccountData struct {
	ID int64
	Username string
	RegistrationDateTime sql.

}

type ModuleData struct {
	ID int64
	Name string
	ParentModuleID sql.NullInt64
}

type Module interface {
	Create(moduleData *ModuleData) error //missing model data
	GetByID(id uint64) (error, *ModuleData)
	DeleteByID(id uint64) error
	Update(ModuleData *ModuleData) error
}

type ModuleData struct {
	ID int64
	Name string
	ParentModuleID sql.NullInt64
}


type OperationData struct {
	ID int64
	Name string
	ModuleID int64
}

type Operation interface {
	Create(operationData *OperationData) error
	GetByID(id int64) (error, *OperationData)
	DeleteByID(id int64) error
	Update(OperationData *OperationData) error
}

type PermissionData struct {
	ID int64
	Description string
}

type Permission interface {
	Create(permissionData *PermissionData) error
	GetByID(id int64) (error, *PermissionData)
	DeleteByID(id int64) error
	Update(permissionData *PermissionData) error
}


type RoleData struct {
	ID int64
	Description    string
	ParentRoleID sql.NullInt64
}

type Role interface {
	Create (roleData *RoleData) error
	GetByID(id int64) (error, *RoleData)
	DeleteByID(id int64) error
	Update(roleData *RoleData) error
}


type OperationPermissionManager interface {
	SetPermissionOperations(permissionID int64, operationIDs []int64) error
	AddPermissionOperations(permissionID int64, operationIDs  []int64) error
	RemovePermissionOperations(permissionID int64, operationIDs  []int64) error
	GetOperationPermissions(operationID int64) (error, []*PermissionData)
	GetPermissionOperations(permissionID int64) (error, []*OperationData)
}


type PermissionRoleManager interface {
	SetRolePermissions(roleID int64, permissionsIDs []int64) error
	AddRolePermissions(roleID int64, permissionsIDs []int64) error
	RemoveRolePermissions(roleID int64, permissionsIDs []int64) error
	GetPermissionRoles(permissionID int64) (error, []*RoleData)
	GetRolePermissions(roleID int64) (error, []*PermissionData)
}


