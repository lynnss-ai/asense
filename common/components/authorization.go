// Package components
// @File    : authorization.go
// @Author  : Wang Xuebing
// @Contact : lynnss.ai@hotmail.com
// @Time    : 2024/9/9 16:57
// @Desc    :
package components

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
)

const (
	CtxAuth = "Authorization"
)

const (
	CtxAuthKeyByJWTUserID = "JwtUserID"
)

func GeneratorJwtToken(val string, secretKey string, iat, seconds int64, userId string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["iss"] = "lynnss.ai@hotmail.com"
	claims[val] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func GetAuthJwtKeyValue[T any](ctx context.Context, val string) T {
	value := ctx.Value(val)
	return value.(T)
}

func GetAuthKeyJwtUserID(ctx context.Context) string {
	return GetAuthJwtKeyValue[string](ctx, CtxAuthKeyByJWTUserID)
}
