// Copyright Letga Author(https://letga.net). All Rights Reserved.
// Apache License 2.0(https://github.com/lgcgo/letga-server/blob/main/LICENSE)

package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 声明格式
// RegisteredClaims 包含了JWT给出的7个官方字段
// - iss (issuer)：发布者，通常填域名即可
// - sub (subject)：签发主题，一般用使用用户的唯一标识
// - iat (Issued At)：生成签名的时间
// - exp (expiration time)：签名过期时间
// - aud (audience)：观众，相当于接受者（客户端/浏览器）
// - nbf (Not Before)：生效时间
// - jti (JWT ID)：编号
type CustomClaims struct {
	IssueType string `json:"ist"` // 签发类型, grant=授予,renew=刷新
	IssueRole string `json:"isr"` // 签发角色, 签发的角色名称（允许多角色）
	jwt.RegisteredClaims
}

type IssueClaims struct {
	Type    string
	Role    string
	Subject string
}

// 配置项目，选项模式
type Config struct {
	Issuer   string
	Audience []string
	SignKey  []byte
}

var (
	config *Config
)

func Init(cfg *Config) {
	config = cfg
}

// 签发Token
// https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-NewWithClaims-CustomClaimsType
func NewToken(issueClaims IssueClaims, expireTime time.Duration) (string, error) {
	claims := CustomClaims{
		issueClaims.Type,
		issueClaims.Role,
		jwt.RegisteredClaims{
			Issuer:    config.Issuer,
			Subject:   issueClaims.Subject,
			Audience:  config.Audience,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(config.SignKey)
}

// 解析Token
// https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-Parse-Hmac
func ParseToken(tokenString string) (map[string]interface{}, error) {
	var (
		token  *jwt.Token
		claims map[string]interface{}
		err    error
		ok     bool
	)

	// 解析Token
	if token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok = token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.SignKey, nil
	}); err != nil {
		return nil, err
	}
	// 校验Token
	if !token.Valid {
		return nil, errors.New("token valid fail")
	}
	// 格式转换
	if claims, ok = token.Claims.(jwt.MapClaims); !ok {
		return nil, errors.New("mapclaims types error")
	}

	return claims, nil
}
