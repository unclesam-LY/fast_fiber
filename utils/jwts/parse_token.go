package jwts

import (
	"FastFiber/global"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

// CheckAccessToken 验证访问令牌
func CheckAccessToken(token string) (*MyClaims, error) {
	return parseToken(token, global.Config.Jwt.AccessSecret)
}

// CheckRefreshToken 验证刷新令牌
func CheckRefreshToken(token string) (*MyClaims, error) {
	return parseToken(token, global.Config.Jwt.RefreshSecret)
}

// parseToken 验证token
func parseToken(token, secret string) (*MyClaims, error) {
	//解析、验证并返回token。
	tokenObj, err := jwt.ParseWithClaims(token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenObj.Claims.(*MyClaims); ok && tokenObj.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token无效")
	}
}
