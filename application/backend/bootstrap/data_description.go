package bootstrap

import "github.com/daniloanp/Ensaios/application/backend/model"


//(map[string][]string)

type module__ struct {
	Module *model.ModuleData
}

var (
	baseModule = &module{
		ModuleData: model.ModuleData{Name: "/"},
		Operations:[]*operation{
			{model.OperationData{Name: ""}, nil}, // base OP, TODO:Recheck
			{model.OperationData{Name: "_login"}, nil},
		},

	}
	adminModule = &module{
		ParentModule: baseModule,
		ModuleData: model.ModuleData{Name: "admin"},
		Operations:nil,

	}
	adminSubModules = []*module{
		{
			ParentModule: adminModule,
			ModuleData: model.ModuleData{Name: "users"},
			Operations:[]*operation{
				{model.OperationData{Name: "listing"}, nil},
				{model.OperationData{Name: "create"}, nil},
				{model.OperationData{Name: "read"}, nil},
				{model.OperationData{Name: "update"}, nil},
			},
		},
		{
			ParentModule: adminModule,
			ModuleData: model.ModuleData{Name: "modules"},
			Operations:[]*operation{
				{model.OperationData{Name: "listing"}, nil},
				{model.OperationData{Name: "create"}, nil},
				{model.OperationData{Name: "read"}, nil},
				{model.OperationData{Name: "update"}, nil},
				{model.OperationData{Name: "delete"}, nil},
				{model.OperationData{Name: "create-operation"}, nil},
				{model.OperationData{Name: "delete-operation"}, nil},
				{model.OperationData{Name: "update-operation"}, nil},
			},
		},
		{
			ParentModule: adminModule,
			ModuleData: model.ModuleData{Name: "modules"},
			Operations:[]*operation{
				{model.OperationData{Name: "listing"}, nil},
				{model.OperationData{Name: "create"}, nil},
				{model.OperationData{Name: "read"}, nil},
				{model.OperationData{Name: "update"}, nil},
				{model.OperationData{Name: "delete"}, nil},
				{model.OperationData{Name: "create-operation"}, nil},
				{model.OperationData{Name: "delete-operation"}, nil},
				{model.OperationData{Name: "update-operation"}, nil},
			},
		},
		{
			ParentModule: adminModule,
			ModuleData: model.ModuleData{Name: "permissions"},
			Operations:[]*operation{
				{model.OperationData{Name: "listing"}, nil},
				{model.OperationData{Name: "create"}, nil},
				{model.OperationData{Name: "read"}, nil},
				{model.OperationData{Name: "update"}, nil},
				{model.OperationData{Name: "delete"}, nil},
				{model.OperationData{Name: "add-operation"}, nil},
				{model.OperationData{Name: "remove-operation"}, nil},
				{model.OperationData{Name: "add-module"}, nil},
				{model.OperationData{Name: "remove-module"}, nil},
			},
		},
		{
			ParentModule: adminModule,
			ModuleData: model.ModuleData{Name: "roles"},
			Operations:[]*operation{
				{model.OperationData{Name: "listing"}, nil},
				{model.OperationData{Name: "create"}, nil},
				{model.OperationData{Name: "read"}, nil},
				{model.OperationData{Name: "update"}, nil},
				{model.OperationData{Name: "delete"}, nil},
				{model.OperationData{Name: "add-permission"}, nil},
			},
		},
	}
)
