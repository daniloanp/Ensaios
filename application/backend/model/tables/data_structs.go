package tables

import (
	"database/sql"
	"time"
)

type (
	UserAccountData struct {
		Id                   int64
		Username             string
		RegistrationDateTime time.Time
	}

	UserPasswordData struct {
		Id                   int64
		UserAccountId        int64
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
		Id                   int64
		UserAccountId        int64
		GivenName            string
		LastName             string
		MotherName           string
		FatherName           string
		Nationality          string
		RegistrationDatetime time.Time
	}

	ModuleData struct {
		Id             int64
		Name           string
		ParentModuleId sql.NullInt64
	}

	OperationData struct {
		Id       int64
		Name     string
		ModuleId int64
	}


	PermissionData struct {
		Id          int64
		Description string
	}


	RoleData struct {
		Id           int64
		Description  string
		ParentRoleId sql.NullInt64
	}
)
