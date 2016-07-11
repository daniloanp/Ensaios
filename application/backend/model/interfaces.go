package model

type(
	UserAccount interface {
		Create(data *UserAccountData) (err error)
		GetByUsername(username string) (*UserAccountData, error)
		GetByID(id int64) (*UserAccountData, error)
		DeleteByID(id int64) (err error)
		Update(data *UserAccountData) (err error)
	}

	UserPassword interface {
		Create(data *UserPasswordData) (err error)
		Update(data *UserPasswordData) (err error)
		DeleteByID(id int64) (err error)
		GetByID(id int64) (*UserPasswordData, error)
	}

	UserEmail interface {
		Create(data *UserEmailData) (err error)
		Update(data *UserEmailData) (err error)
		DeleteByAddress(address string) (err error)
		GetByAddress(address string) (*UserEmailData, error)
	}

	UserPersonalInformation interface {
		Create(userPersonalInformation *UserPersonalInformation) (err error)
		Update(userPersonalInformation  *UserPersonalInformation) (err error)
		DeleteById(id int64) (err error)
		GetById(id int64) (*UserPersonalInformation, error)
	}
	Module interface {
		Create(data *ModuleData) (err error)
		GetByID(id int64) (*ModuleData, error)
		DeleteByID(id int64) (err error)
		Update(data *ModuleData) (err error)
	}
	Operation interface {
		Create(data *OperationData) (err error)
		GetByID(id int64) (*OperationData, error)
		DeleteByID(id int64) (err error)
		Update(data *OperationData) (err error)
	}

	Permission interface {
		Create(data *PermissionData) (err error)
		GetByID(id int64) (*PermissionData, error)
		DeleteByID(id int64) (err error)
		Update(data *PermissionData) (err error)
	}

	Role interface {
		Create(data *RoleData) (err error)
		GetByID(id int64) (*RoleData, error)
		DeleteByID(id int64) (err error)
		Update(data *RoleData) (err error)
		HasPermission(data *RoleData, operationPath string) (bool, error)
	}

	OperationPermissionManager interface {
		SetPermissionOperations(permissionID int64, operationIDs []int64) (err error)
		AddPermissionOperations(permissionID int64, operationIDs  []int64) (err error)
		RemovePermissionOperations(permissionID int64, operationIDs  []int64) (err error)
		GetOperationPermissions(operationID int64) ([]*PermissionData, error)
		GetPermissionOperations(permissionID int64) ([]*OperationData, error)
	}

	PermissionRoleManager interface {
		SetRolePermissions(roleID int64, permissionsIDs []int64) (err error)
		AddRolePermissions(roleID int64, permissionsIDs []int64) (err error)
		RemoveRolePermissions(roleID int64, permissionsIDs []int64) (err error)
		GetPermissionRoles(permissionID int64) ([]*RoleData, error)
		GetRolePermissions(roleID int64) ([]*PermissionData, error)
	}
)

