package server

const (
	AdminPrefix = "_admin/"
)
var adminActions  = map[string]handler {
	//users
	AdminPrefix + "users/listing": nil,
	AdminPrefix + "users/create": nil,
	AdminPrefix + "users/read": nil,
	AdminPrefix + "users/update": nil,
	//modules
	AdminPrefix + "modules/listing": nil,
	AdminPrefix + "modules/create" : nil,
	AdminPrefix + "modules/read" : nil,
	AdminPrefix + "modules/update" : nil,
	AdminPrefix + "modules/delete" : nil,
	//modules-operation
	AdminPrefix + "modules/create-operation" : nil,
	AdminPrefix + "modules/delete-operation" : nil,
	AdminPrefix + "modules/update-operation" : nil,

	//Permissions
	AdminPrefix + "permissions/listing": nil,
	AdminPrefix + "permissions/read": nil,
	AdminPrefix + "permissions/update": nil,
	AdminPrefix + "permissions/create": nil,
	AdminPrefix + "permissions/delete": nil,

	AdminPrefix + "permissions/add-operation": nil,
	AdminPrefix + "permissions/remove-operation": nil,
	AdminPrefix + "permissions/add-module": nil, // Just a shortcut
	AdminPrefix + "permissions/remove-module": nil, // Just a shortcut

	//Roles
	AdminPrefix + "roles/listing": nil,
	AdminPrefix + "roles/read": nil,
	AdminPrefix + "roles/update": nil,
	AdminPrefix + "roles/create": nil,
	AdminPrefix + "roles/delete": nil,
	AdminPrefix + "roles/add-permission": nil,
}

