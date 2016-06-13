package app

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"database/sql"
)



func createAdminModule (baseModuleID int64) {
	var err error
	var adminModule = &model.ModuleData{
		Name:"_admin",
		ParentModuleID: sql.NullInt64{
			Int64: baseModuleID,
			Valid: true,
		},
	}
	err = db.Module.Create(adminModule)
	throwPanic(err)

	var usersModule = &model.ModuleData{
		Name:"users",
		ParentModuleID: sql.NullInt64{
			Int64: adminModule.ID,
			Valid: true,
		},
	}

	err = db.Module.Create(usersModule)
	throwPanic(err)

	var operationsModule = &model.ModuleData {
		Name:"operations",
		ParentModuleID: sql.NullInt64{
			Int64: adminModule.ID,
			Valid: true,
		},
	}

	err = db.Module.Create(operationsModule)
	throwPanic(err)

	var permissionsModule = &model.ModuleData {
		Name: "permissions",
		ParentModuleID: sql.NullInt64{
			Int64: adminModule.ID,
			Valid: true,
		},
	}
	err = db.Module.Create(permissionsModule)
	throwPanic(err)

	var rolesModule = &model.ModuleData {
		Name: "roles",
		ParentModuleID: sql.NullInt64{
			Int64: adminModule.ID,
			Valid: true,
		},
	}

	err = db.Module.Create(rolesModule)
	throwPanic(err)
}



func createBaseModule () *model.ModuleData {
	var err error
	var baseModule = &model.ModuleData{
		Name:"",
	}
	err = db.Module.Create(baseModule)
	throwPanic(err)

	var loginOp = &model.OperationData{
		Name:"_login",
		ModuleID: baseModule.ID,
	}

	err = db.Operation.Create(loginOp)
	throwPanic(err)

	var baseOp = &model.OperationData{
		Name:"",
		ModuleID: baseModule.ID,
	}

	err = db.Operation.Create(baseOp)
	throwPanic(err)

	return baseModule;
}


func Bootstrap() {
	baseModule := createBaseModule()
	createAdminModule(baseModule.ID)
}
