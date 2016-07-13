package bootstrap

import (
	"github.com/daniloanp/Ensaios/application/backend/model/tables"
)

type module__ struct {
	Module *tables.ModuleData
}

var (
	baseModule = &module{
		ModuleData: tables.ModuleData{Name: "/"},
		Operations:[]*operation{
			{tables.OperationData{Name: ""}, nil}, // base OP, TODO:Recheck
			{tables.OperationData{Name: "_login"}, nil},
		},
	}
	adminModule = &module{
		ParentModule: baseModule,
		ModuleData: tables.ModuleData{Name: "admin"},
		Operations:nil,

	}
	adminSubModules = []*module{
		{
			ParentModule: adminModule,
			ModuleData: tables.ModuleData{Name: "users"},
			Operations:[]*operation{
				{tables.OperationData{Name: "listing"}, nil},
				{tables.OperationData{Name: "create"}, nil},
				{tables.OperationData{Name: "read"}, nil},
				{tables.OperationData{Name: "update"}, nil},
			},
		},
		{
			ParentModule: adminModule,
			ModuleData: tables.ModuleData{Name: "modules"},
			Operations:[]*operation{
				{tables.OperationData{Name: "listing"}, nil},
				{tables.OperationData{Name: "create"}, nil},
				{tables.OperationData{Name: "read"}, nil},
				{tables.OperationData{Name: "update"}, nil},
				{tables.OperationData{Name: "delete"}, nil},
				{tables.OperationData{Name: "create-operation"}, nil},
				{tables.OperationData{Name: "delete-operation"}, nil},
				{tables.OperationData{Name: "update-operation"}, nil},
			},
		},
		{
			ParentModule: adminModule,
			ModuleData: tables.ModuleData{Name: "modules"},
			Operations:[]*operation{
				{tables.OperationData{Name: "listing"}, nil},
				{tables.OperationData{Name: "create"}, nil},
				{tables.OperationData{Name: "read"}, nil},
				{tables.OperationData{Name: "update"}, nil},
				{tables.OperationData{Name: "delete"}, nil},
				{tables.OperationData{Name: "create-operation"}, nil},
				{tables.OperationData{Name: "delete-operation"}, nil},
				{tables.OperationData{Name: "update-operation"}, nil},
			},
		},
		{
			ParentModule: adminModule,
			ModuleData: tables.ModuleData{Name: "permissions"},
			Operations:[]*operation{
				{tables.OperationData{Name: "listing"}, nil},
				{tables.OperationData{Name: "create"}, nil},
				{tables.OperationData{Name: "read"}, nil},
				{tables.OperationData{Name: "update"}, nil},
				{tables.OperationData{Name: "delete"}, nil},
				{tables.OperationData{Name: "add-operation"}, nil},
				{tables.OperationData{Name: "remove-operation"}, nil},
				{tables.OperationData{Name: "add-module"}, nil},
				{tables.OperationData{Name: "remove-module"}, nil},
			},
		},
		{
			ParentModule: adminModule,
			ModuleData: tables.ModuleData{Name: "roles"},
			Operations:[]*operation{
				{tables.OperationData{Name: "listing"}, nil},
				{tables.OperationData{Name: "create"}, nil},
				{tables.OperationData{Name: "read"}, nil},
				{tables.OperationData{Name: "update"}, nil},
				{tables.OperationData{Name: "delete"}, nil},
				{tables.OperationData{Name: "add-permission"}, nil},
			},
		},
	}
)
