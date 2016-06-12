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

type UserPasswordData struct {
	ID                   int64
	UserAccountID        int64
	Password             string
	Salt                 string
	RegistrationDateTime time.Time
}

type UserEmailData struct {
	UserAccountId        int64
	Address              string
	Verified             bool
	RegistrationDateTime time.Time
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

type ModuleData struct {
	ID             int64
	Name           string
	ParentModuleID sql.NullInt64
}

type OperationData struct {
	ID       int64
	Name     string
	ModuleID int64
}

type PermissionData struct {
	ID          int64
	Description string
}

type RoleData struct {
	ID           int64
	Description  string
	ParentRoleID sql.NullInt64
}
