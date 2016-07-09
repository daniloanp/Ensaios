package model

type(
	UserAccount interface {
		Create(userAccountData *UserAccountData) error
		GetByUsername(username string) (*UserAccountData, error)
		GetByID(id int64) (*UserAccountData, error)
		DeleteByID(id int64) error
		Update(userAccountData *UserAccountData) error
	}

	UserPassword interface {
		Create(userPasswordData *UserPasswordData) error
		Update(userPasswordData *UserPasswordData) error
		DeleteByID(id int64) error
		GetByID(id int64) (*UserPasswordData, error)
	}

	UserEmail interface {
		Create(userEmailData *UserEmailData) error
		Update(userEmailData  *UserEmailData) error
		DeleteByAddress(address string) error
		GetByAddress(address string) (*UserEmailData, error)
	}

	UserPersonalInformation interface {
		Create(userPersonalInformation *UserPersonalInformation) error
		Update(userPersonalInformation  *UserPersonalInformation) error
		DeleteById(id int64) error
		GetById(id int64) (*UserPersonalInformation, error)
	}
	Module interface {
		Create(moduleData *ModuleData) error //missing model data
		GetByID(id int64) (*ModuleData, error)
		DeleteByID(id int64) error
		Update(ModuleData *ModuleData) error
	}
	Operation interface {
		Create(operationData *OperationData) error
		GetByID(id int64) (*OperationData, error)
		DeleteByID(id int64) error
		Update(OperationData *OperationData) error
	}

	Permission interface {
		Create(permissionData *PermissionData) error
		GetByID(id int64) (*PermissionData, error)
		DeleteByID(id int64) error
		Update(permissionData *PermissionData) error
	}

	Role interface {
		Create(roleData *RoleData) error
		GetByID(id int64) (*RoleData, error)
		DeleteByID(id int64) error
		Update(roleData *RoleData) error
		HasPermission(roleData *RoleData, operationPath string) (bool, error) // TODO:Maybe we should not tell when a op exists
	}

	OperationPermissionManager interface {
		SetPermissionOperations(permissionID int64, operationIDs []int64) error
		AddPermissionOperations(permissionID int64, operationIDs  []int64) error
		RemovePermissionOperations(permissionID int64, operationIDs  []int64) error
		GetOperationPermissions(operationID int64) ([]*PermissionData, error)
		GetPermissionOperations(permissionID int64) ([]*OperationData, error)
	}

	PermissionRoleManager interface {
		SetRolePermissions(roleID int64, permissionsIDs []int64) error
		AddRolePermissions(roleID int64, permissionsIDs []int64) error
		RemoveRolePermissions(roleID int64, permissionsIDs []int64) error
		GetPermissionRoles(permissionID int64) ([]*RoleData, error)
		GetRolePermissions(roleID int64) ([]*PermissionData, error)
	}
)

