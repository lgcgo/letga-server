// Copyright Letga Author(https://letga.net). All Rights Reserved.
// Apache License 2.0(https://github.com/lgcgo/letga-server/blob/main/LICENSE)

package auth

import (
	"context"
	"letga/internal/biz"
	"letga/internal/dao"
	"letga/internal/model"
	"letga/internal/service"
	"letga/utility/casbin"
	"letga/utility/token"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type Config struct {
	TokenType              string
	AccessTokenExpireTime  time.Duration
	RefreshTokenExpireTime time.Duration
}

type sAuth struct {
	Config *Config
}

func init() {
	service.RegisterAuth(New())
}

func New() service.IAuth {
	var (
		ctx                       = gctx.New()
		accessTokenExpireTimeCfg  = g.Cfg().MustGet(ctx, "token.accessToken.expireTime").String()
		refreshTokenExpireTimeCfg = g.Cfg().MustGet(ctx, "token.refreshToken.expireTime").String()
		tokenTypeCfg              = g.Cfg().MustGet(ctx, "token.tokenType").String()
		accessTokenExpireTime     time.Duration
		refreshTokenExpireTime    time.Duration
		err                       error
	)
	if accessTokenExpireTime, err = time.ParseDuration(accessTokenExpireTimeCfg); err != nil {
		panic("Missing config: token.accessToken.expireTime")
	}
	if refreshTokenExpireTime, err = time.ParseDuration(refreshTokenExpireTimeCfg); err != nil {
		panic("Missing config: token.refreshToken.expireTime")
	}
	return &sAuth{
		Config: &Config{
			TokenType:              tokenTypeCfg,
			AccessTokenExpireTime:  accessTokenExpireTime,
			RefreshTokenExpireTime: refreshTokenExpireTime,
		},
	}
}

// 授权, 支持Bearer签发方案(Header Authorization: Bearer <token>)
func (s *sAuth) Authorization(ctx context.Context, in *model.AuthAuthorizationInput) (*model.AuthToken, error) {
	var (
		currentTime  = time.Now()
		out          *model.AuthToken
		err          error
		accessToken  string
		refreshToken string
		expiresIn    float64
	)
	// 制作AccessToken
	if accessToken, err = token.NewToken(token.IssueClaims{
		Type:    "grant",
		Role:    in.Role,
		Subject: in.Subject,
	}, s.Config.AccessTokenExpireTime); err != nil {
		return nil, gerror.NewCode(biz.AuthTokenIssueFail)
	}
	// 制作RefreshToken
	if refreshToken, err = token.NewToken(token.IssueClaims{
		Type:    "renew",
		Role:    in.Role,
		Subject: in.Subject,
	}, s.Config.RefreshTokenExpireTime); err != nil {
		return nil, gerror.NewCode(biz.AuthTokenIssueFail)
	}
	// 获取过期秒数
	expiresIn = currentTime.Add(s.Config.AccessTokenExpireTime).Sub(currentTime).Seconds()
	out = &model.AuthToken{
		AccessToken:  accessToken,
		TokenType:    s.Config.TokenType,
		RefreshToken: refreshToken,
		ExpiresIn:    uint(expiresIn),
	}
	return out, nil
}

// 刷新授权
func (s *sAuth) RefreshAuthorization(ctx context.Context, tokenString string) (*model.AuthToken, error) {
	var (
		claims map[string]interface{}
		err    error
	)
	// 解析token
	if claims, err = token.ParseToken(tokenString); err != nil {
		return nil, err
	}
	// 校验签发类型
	if claims["ist"] != "renew" {
		return nil, gerror.NewCode(biz.AuthTokenInvalid)
	}
	return s.Authorization(ctx, &model.AuthAuthorizationInput{
		Subject: claims["sub"].(string),
		Role:    claims["isr"].(string),
	})
}

// 验证Token
func (s *sAuth) VerifyToken(ctx context.Context, tokenString string) (map[string]interface{}, error) {
	var (
		claims map[string]interface{}
		err    error
	)
	// 解析Token
	if claims, err = token.ParseToken(tokenString); err != nil {
		return nil, gerror.NewCode(biz.AuthTokenInvalid)
	}
	// 非法的签发类型
	if claims["ist"] != "grant" {
		return nil, gerror.NewCode(biz.AuthTokenInvalid)
	}
	return claims, nil
}

// 验证路由
func (s *sAuth) Verify(ctx context.Context, in *model.AuthVerifyInput) error {
	var (
		ok  bool
		err error
	)
	// 转换入参
	if ok, err = casbin.Verify(&casbin.VerifyPayload{
		Subject: in.Subject,
		Role:    in.Role,
		Path:    in.Path,
		Method:  in.Method,
	}); err != nil {
		return err
	}
	if !ok {
		return gerror.NewCode(biz.AuthNotPermission)
	}

	return nil
}

// 更新授权政策
func (s *sAuth) RefreshPolicy(ctx context.Context) error {
	var (
		routePolicys []*casbin.RoutePolicy // 路由政策集
		rolePolicys  []*casbin.RolePolicy  // 角色政策集
		// userPolicys  []*casbin.UserPolicy  // 用户授权政策集
		roles      []*model.AuthRole
		roleAccess []*model.AuthRoleAccess
		err        error
	)
	// 获取所有角色
	if roles, err = dao.AuthRole.GetAll(ctx); err != nil {
		return err
	}
	// 组装角色政策
	for _, role := range roles {
		var (
			parentRole     *model.AuthRole
			parentRoleName string
		)
		// 过滤超级管理员
		if role.Name == casbin.ROOT_ROLE_NAME {
			continue
		}
		if role.ParentId == 0 {
			// 超级管理员
			parentRoleName = casbin.ROOT_ROLE_NAME
		} else {
			if parentRole, err = s.GetRole(ctx, role.ParentKey); err != nil {
				return err
			}
			parentRoleName = parentRole.Name
		}
		rolePolicys = append(rolePolicys, &casbin.RolePolicy{
			ParentRole: parentRoleName,
			Role:       role.Name,
		})
	}

	// 获取所有角色路由
	var (
		role  *model.AuthRole
		route *model.AuthRoute
	)
	if roleAccess, err = dao.AuthRoleAccess.GetAll(ctx); err != nil {
		return err
	}
	for _, v := range roleAccess {
		if role, err = dao.AuthRole.Get(ctx, v.RoleId); err != nil {
			return err
		}
		if route, err = dao.AuthRoute.Get(ctx, v.RouteId); err != nil {
			return err
		}
		if route == nil {
			continue
		}
		routePolicys = append(routePolicys, &casbin.RoutePolicy{
			Role:   role.Name,
			Path:   route.Path,
			Method: route.Method,
		})
	}
	return casbin.SavePolicyCsv(&casbin.CsvPlicyPaylod{
		RoutePolicys: routePolicys,
		RolePolicys:  rolePolicys,
		// UserPolicys:  userPolicys,
	})
}
