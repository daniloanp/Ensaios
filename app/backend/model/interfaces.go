package model

import (
	"database/sql"
	"time"
)

type UserAccountData struct {
	ID                   int64
	Username             string
	RegistrationDateTime time.Time
}

type UserAccount interface {
	Create(userAccountData *UserAccountData) error
	GetByUsername(username string) (error, *UserAccountData)
	GetByID(id int64) (error, *UserAccountData)
	DeleteByID(id int64) error
	Update(userAccountData *UserAccountData) error
}

type UserPasswordData struct {
	ID                   int64
	UserAccountID        int64
	Password             string
	Salt                 string
	RegistrationDateTime time.Time
}

type UserPassword interface {
	Create(userPasswordData *UserPasswordData) error
	Update(userPasswordData *UserPasswordData) error
	DeleteByID(id int64) error
	GetByID(id int64) (error, *UserPasswordData)
}


type UserEmailData struct {
	UserAccountId        int64
	Address              string
	Verified             bool
	RegistrationDateTime time.Time
}

type UserEmail interface {
	Create(userEmailData *UserEmailData) error
	Update(userEmailData  *UserEmailData) error
	DeleteByAddress(address string) error
	GetByAddress(address string) (error, *UserEmailData)
}

type UserPersonalInformationData struct {
	ID                   int64
	UserAccountID        int64
	GivenName            string
	LastName             string
	MotherName           string
	FatherName           string
	Nationality          string
	RegistrationDatetime time.Time
}

type UserPersonalInformation interface {
	Create(userPersonalInformation *UserPersonalInformation) error
	Update(userPersonalInformation  *UserPersonalInformation) error
	DeleteById(id int64) error
	GetById(id int64) (error, *UserPersonalInformation)
}

type ModuleData struct {
	ID             int64
	Name           string
	ParentModuleID sql.NullInt64
}

type Module interface {
	Create(moduleData *ModuleData) error //missing model data
	GetByID(id uint64) (error, *ModuleData)
	DeleteByID(id uint64) error
	Update(ModuleData *ModuleData) error
}

type OperationData struct {
	ID       int64
	Name     string
	ModuleID int64
}

type Operation interface {
	Create(operationData *OperationData) error
	GetByID(id int64) (error, *OperationData)
	DeleteByID(id int64) error
	Update(OperationData *OperationData) error
}

type PermissionData struct {
	ID          int64
	Description string
}

type Permission interface {
	Create(permissionData *PermissionData) error
	GetByID(id int64) (error, *PermissionData)
	DeleteByID(id int64) error
	Update(permissionData *PermissionData) error
}

type RoleData struct {
	ID           int64
	Description  string
	ParentRoleID sql.NullInt64
}

type Role interface {
	Create(roleData *RoleData) error
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


