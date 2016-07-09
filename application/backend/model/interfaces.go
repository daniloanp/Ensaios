package model

type(
	UserAccount interface {
		Create(userAccountData *UserAccountData) error
		GetByUsername(username string) (error, *UserAccountData)
		GetByID(id int64) (error, *UserAccountData)
		DeleteByID(id int64) error
		Update(userAccountData *UserAccountData) error
	}
	UserPassword interface {
		Create(userPasswordData *UserPasswordData) error
		Update(userPasswordData *UserPasswordData) error
		DeleteByID(id int64) error
		GetByID(id int64) (error, *UserPasswordData)
	}
	UserEmail interface {
		Create(userEmailData *UserEmailData) error
		Update(userEmailData  *UserEmailData) error
		DeleteByAddress(address string) error
		GetByAddress(address string) (error, *UserEmailData)
	}
	UserPersonalInformation interface {
		Create(userPersonalInformation *UserPersonalInformation) error
		Update(userPersonalInformation  *UserPersonalInformation) error
		DeleteById(id int64) error
		GetById(id int64) (error, *UserPersonalInformation)
	}
	Module interface {
		Create(moduleData *ModuleData) error //missing model data
		GetByID(id uint64) (error, *ModuleData)
		DeleteByID(id uint64) error
		Update(ModuleData *ModuleData) error
	}
	Operation interface {
		Create(operationData *OperationData) error
		GetByID(id int64) (error, *OperationData)
		DeleteByID(id int64) error
		Update(OperationData *OperationData) error
	}
	Permission interface {
		Create(permissionData *PermissionData) error
		GetByID(id int64) (error, *PermissionData)
		DeleteByID(id int64) error
		Update(permissionData *PermissionData) error
	}
	Role interface {
		Create(roleData *RoleData) error
		GetByID(id int64) (error, *RoleData)
		DeleteByID(id int64) error
		Update(roleData *RoleData) error
		HasPermission(roleData *RoleData, operationPath string) (error, bool) // TODO:Maybe we should not tell when a op exists
	}
	OperationPermissionManager interface {
		SetPermissionOperations(permissionID int64, operationIDs []int64) error
		AddPermissionOperations(permissionID int64, operationIDs  []int64) error
		RemovePermissionOperations(permissionID int64, operationIDs  []int64) error
		GetOperationPermissions(operationID int64) (error, []*PermissionData)
		GetPermissionOperations(permissionID int64) (error, []*OperationData)
	}
	PermissionRoleManager interface {
		SetRolePermissions(roleID int64, permissionsIDs []int64) error
		AddRolePermissions(roleID int64, permissionsIDs []int64) error
		RemoveRolePermissions(roleID int64, permissionsIDs []int64) error
		GetPermissionRoles(permissionID int64) (error, []*RoleData)
		GetRolePermissions(roleID int64) (error, []*PermissionData)
	}
)

