package tables

import "github.com/daniloanp/Ensaios/application/backend/model"

var tbMap *tableMap = nil

func Instance() model.Tb {
	if tbMap == nil {
		tbMap = &tableMap{
			userAccount: new(userAccount),
			userPassword: new(userPassword),
			userEmail: new(userEmail),
			userPersonalInformation: new(userPersonalInformation),
			module: new(module),
			operation: new(operation),
			permission: new(permission),
			role: new(role),
		}
	}
	return tbMap
}


