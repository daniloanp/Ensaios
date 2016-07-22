package tables

type(
	UserAccount interface {
		Create(data *UserAccountData) (err error)
		GetByUsername(username string) (*UserAccountData, error)
		GetById(id int64) (*UserAccountData, error)
		DeleteById(id int64) (err error)
		Update(data *UserAccountData) (err error)
	}

	UserPassword interface {
		Create(data *UserPasswordData) (err error)
		Update(data *UserPasswordData) (err error)
		DeleteById(id int64) (err error)
		GetById(id int64) (*UserPasswordData, error)
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
		GetById(id int64) (*ModuleData, error)
		DeleteById(id int64) (err error)
		Update(data *ModuleData) (err error)
	}
	Operation interface {
		Create(data *OperationData) (err error)
		GetById(id int64) (*OperationData, error)
		DeleteById(id int64) (err error)
		Update(data *OperationData) (err error)
	}

	Permission interface {
		Create(data *PermissionData) (err error)
		GetById(id int64) (*PermissionData, error)
		DeleteById(id int64) (err error)
		Update(data *PermissionData) (err error)
	}

	Role interface {
		Create(data *RoleData) (err error)
		GetById(id int64) (*RoleData, error)
		DeleteById(id int64) (err error)
		Update(data *RoleData) (err error)
		HasPermission(data *RoleData, operationPath string) (bool, error)
	}
)

