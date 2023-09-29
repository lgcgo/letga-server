// Copyright Letga Author(https://letga.net). All Rights Reserved.
// Apache License 2.0(https://github.com/lgcgo/letga-server/blob/main/LICENSE)

package biz

var (
	/**
	 * 公共错误码
	 */
	// TableIdDecodeFail = WithCode(CodeOk, 11001, "table id '%s' decode fail")
	TableKeyInvalid = WithCode(CodeOk, 11002, "table key '%s' invalid")
	// TableKeysInvalid = WithCode(CodeOk, 11002, "table keys '%s' invalid")

	TreeInitialFail = WithCode(CodeOk, 12001, "tree initial fail")

	/*
	 * 授权服务
	 * 错误码:10000-19999
	 */
	// 权限认证
	AuthHeaderInvalid = WithCode(CodeOk, 10001, "header token invalid")
	AuthTokenInvalid  = WithCode(CodeOk, 10002, "token invalid")
	AuthNotPermission = WithCode(CodeOk, 10003, "no resource access permission")

	// 访问授权
	AuthTokenIssueFail = WithCode(CodeOk, 11001, "token issue failed")
	AuthCasbinInitFail = WithCode(CodeOk, 11002, "casbin initialization failed")

	// 用户授权
	AuthAccessNotExists   = WithCode(CodeOk, 22001, "access does not exists")
	AuthAccessRoleInvalid = WithCode(CodeOk, 16002, "access role '%s' invalid")

	// 权限角色
	AuthRoleNotExists       = WithCode(CodeOk, 14001, "role does not exists")
	AuthRoleInvalid         = WithCode(CodeOk, 14002, "role invalid")
	AuthRoleNameInvalid     = WithCode(CodeOk, 14002, "role's name invalid")
	AuthRoleNameExists      = WithCode(CodeOk, 14003, "role's name is already exists")
	AuthRoleIdsInvalid      = WithCode(CodeOk, 14004, "role ids '%s' invalid")
	AuthRoleParentInvalid   = WithCode(CodeOk, 14005, "role parent invalid")
	AuthRoleParentNotExists = WithCode(CodeOk, 14006, "role parent does not exists")
	AuthRoleAccessInvalid   = WithCode(CodeOk, 14007, "role access invalid")

	// 角色授权
	AuthRoleAccessExists    = WithCode(CodeOk, 14004, "role access is exists")
	AuthRoleAccessNotExists = WithCode(CodeOk, 15002, "role access not exists")

	// 权限路由
	AuthRouteNotExists = WithCode(CodeOk, 16001, "route not exists.")
	AuthRouteExists    = WithCode(CodeOk, 16002, "route '%s' is already exists.")
	AuthRouteInvalid   = WithCode(CodeOk, 16003, "route invalid")
	// AuthRouteIdsInvalid = WithCode(CodeOk, 16003, "route ids %s invalid.")

	/*
	 * 会员服务
	 * 错误码:20000-29999
	 */
	// 注册与登录
	UserPassportInvalid      = WithCode(CodeOk, 20002, "passport field is invalid")
	UserSignIncorrect        = WithCode(CodeOk, 20003, "passport or password invalid")
	UserSigninTypeIcorrect   = WithCode(CodeOk, 20004, "passport or password invalid")
	UserAccountExists        = WithCode(CodeOk, 20005, "account '%s' is already exists")
	UserAccountNotExists     = WithCode(CodeOk, 20006, "account does not exists")
	UserMobileExists         = WithCode(CodeOk, 20007, "mobile '%s' is already exists")
	UserEmailExists          = WithCode(CodeOk, 20008, "email '%s' is already exists")
	UserPasswordEmpty        = WithCode(CodeOk, 20009, "password cannot be empty")
	UserSigninFailureInvalid = WithCode(CodeOk, 20010, "signin failure limit")
	UserSigninTimeOut        = WithCode(CodeOk, 20011, "signin timed out")

	// 会员操作
	UserNotExists     = WithCode(CodeOk, 21001, "user does not exists")
	UserIdsInvalid    = WithCode(CodeOk, 21003, "user ids '%s' invalid")
	UserStatusInvalid = WithCode(CodeOk, 20001, "user status is invalid")

	// 媒体操作
	MediaNotExists           = WithCode(CodeOk, 21001, "media '%s' does not exists")
	MediaUnknownType         = WithCode(CodeOk, 21001, "unknown media type")
	MediaReadFail            = WithCode(CodeOk, 21001, "media read fail")
	MediaUrlParseFail        = WithCode(CodeOk, 21001, "media url parse fail")
	MediaHashCalculationFail = WithCode(CodeOk, 21001, "media hash calculation fail")
	MediaFileTypeInvalid     = WithCode(CodeOk, 20001, "media file type is invalid")
	MediaMimeTypeInvalid     = WithCode(CodeOk, 20001, "media mime type is invalid")
	MediaIdsInvalid          = WithCode(CodeOk, 12003, "media ids '%s' invalid")

	// 菜单
	MenuNotExists       = WithCode(CodeOk, 13001, "menu does not exists")
	MenuNameExists      = WithCode(CodeOk, 13002, "menu's name '%s' is already exists")
	MenuIdsInvalid      = WithCode(CodeOk, 13003, "menu ids '%s' invalid")
	MenuParentInvalid   = WithCode(CodeOk, 13004, "menu parent invalid")
	MenuParentNotExists = WithCode(CodeOk, 13005, "menu parent does not exists")
)
