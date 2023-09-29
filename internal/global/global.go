package global

import "letga/internal/consts"

/**
* 全局变量
**/

// map{"user": map{id:101}}
var HashidSaltMap = map[string]map[string]int{
	consts.TABLE_AUTH_ROLE: {
		"id":        consts.TABLE_AUTH_ROLE_SALT,
		"parent_id": consts.TABLE_AUTH_ROLE_SALT,
	},
	consts.TABLE_AUTH_ROLE_ACCESS: {
		"id":       consts.TABLE_AUTH_ROLE_ACCESS_SALT,
		"role_id":  consts.TABLE_AUTH_ROLE_SALT,
		"route_id": consts.TABLE_AUTH_ROUTE_SALT,
	},
	consts.TABLE_AUTH_ROUTE: {
		"id":      consts.TABLE_AUTH_ROUTE_SALT,
		"menu_id": consts.TABLE_MENU_SALT,
	},
	consts.TABLE_MEDIA: {
		"id":      consts.TABLE_MEDIA_SALT,
		"user_id": consts.TABLE_USER_SALT,
	},
	consts.TABLE_MENU: {
		"id":        consts.TABLE_MENU_SALT,
		"parent_id": consts.TABLE_MENU_SALT,
	},
	consts.TABLE_USER: {
		"id": consts.TABLE_USER_SALT,
	},
	consts.TABLE_AUTH_ACCESS: {
		"id":      consts.TABLE_AUTH_ACCESS_SALT,
		"user_id": consts.TABLE_USER_SALT,
		"role_id": consts.TABLE_AUTH_ROLE_SALT,
	},
}
