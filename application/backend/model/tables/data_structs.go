package tables

import (
	"database/sql"
	"time"
)

type (
	UserAccountData struct {
		ID                   int64
		Username             string
		RegistrationDateTime time.Time
	}
	UserPasswordData struct {
		ID                   int64
		UserAccountID        int64
		Password             string
		Salt                 string
		RegistrationDateTime time.Time
	}
	UserEmailData struct {
		UserAccountId        int64
		Address              string
		Verified             bool
		RegistrationDateTime time.Time
	}
	UserPersonalInformationData struct {
		ID                   int64
		UserAccountID        int64
		GivenName            string
		LastName             string
		MotherName           string
		FatherName           string
		Nationality          string
		RegistrationDatetime time.Time
	}
	ModuleData struct {
		ID             int64
		Name           string
		ParentModuleID sql.NullInt64
	}
	OperationData struct {
		ID       int64
		Name     string
		ModuleID int64
	}
	PermissionData struct {
		ID          int64
		Description string
	}
	RoleData struct {
		ID           int64
		Description  string
		ParentRoleID sql.NullInt64
	}
)
