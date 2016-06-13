package app

import (
	"github.com/daniloanp/Ensaios/application/backend/model"
	"database/sql"
)

type module struct {
	Module *model.ModuleData
	Operations []*model.OperationData
}

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

func createAdminModule (baseModuleID int64) (*model.ModuleData, map[string]*module) {
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
		createOperations(m.Module.ID, operations)
	}

	return adminModule, subModules

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

func createPermissionForAdmin(adminModuleData *model.ModuleData, subModules map[string]*module) {
	var err error
	//TODO:This should be reactored to splitt_up_permissions
	var adminPermission = &model.PermissionData{
		Description: "Full administration Permission",
	}
	err = db.Operation.Create(adminPermission)
	throwPanic(err)


	for _, module := range subModules {
		operationIds := make([]int64, 0, len(module.Operations))
		for operation := range module.Operations {
			operationIds = append(operationIds, operation)
		}
		db.OperationPermissionManager.SetPermissionOperations(adminPermission.ID)
	}
}


func Bootstrap() {
	baseModule := createBaseModule()
	adminModuleData, subModules := createAdminModule(baseModule.ID)
	createPermissionForAdmin(adminModuleData, subModules)
}
