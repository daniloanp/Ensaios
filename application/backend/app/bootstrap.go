package app

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"database/sql"
)

//(map[string][]string)
var adminSubModules = map[string][]string{
	"users": {
		"listing",
		"create",
		"read",
		"update",
	},
	"modules": {
		"listing",
		"create",
		"read",
		"update",
		"delete",
		"create-operation",
		"delete-operation",
		"update-operation",
	},
	"permissions": {
		"listing",
		"read",
		"update",
		"create",
		"delete",
		"add-operation",
		"remove-operation",
		"add-module",
		"remove-module",
	},
	"roles" : {
		"listing",
		"read",
		"update",
		"create",
		"delete",
		"add-permission",
	},
}



func createModule(moduleName string, parentID int64) *model.ModuleData {
	var err error
	var module = &model.ModuleData{
		Name:moduleName,
		ParentModuleID: sql.NullInt64{
			Int64: parentID,
			Valid: true,
		},
	}

	err = db.Module.Create(module)
	throwPanic(err)

	return module
}

func createOperations(moduleID int64, operationNames []string) []*model.OperationData {
	var err error
	var operationsSlice = make([]*model.OperationData, 0, len(operationNames))
	for _, name := range operationNames {
		var operation = &model.OperationData{
			Name: name,
			ModuleID: moduleID,
		}
		err = db.Operation.Create(operation)
		throwPanic(err)

		operationsSlice = append(operationsSlice, operation)
	}
	return operationsSlice
}

func createAdminModule (baseModuleID int64) {
	type module struct {
		Module *model.ModuleData
		Operations []*model.OperationData
	}
	var err error

	var adminModule = &model.ModuleData{
		Name: "_admin",
		ParentModuleID: sql.NullInt64{
			Int64: baseModuleID,
			Valid: true,
		},
	}
	err = db.Module.Create(adminModule)
	throwPanic(err)

	var subModules = make(map[string]*module, len(adminSubModules))

	for moduleName, operations := range adminSubModules {
		var m = &module{}
		m.Module = createModule(moduleName, adminModule.ID)
		m.Operations = createOperations(
			m.Module.ID,
			operations,
		)
		subModules[moduleName] = m
	}




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
