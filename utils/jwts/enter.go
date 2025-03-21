package jwts

import (
	"FastFiber/global"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MyClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// GenAccessToken 生成Access Token
func GenAccessToken(userID uint) (string, error) {
	claims := MyClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(
				time.Duration(global.Config.Jwt.AccessExpire) * time.Hour)),
			Issuer: global.Config.Jwt.Issuer,
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(global.Config.Jwt.AccessSecret))
}

// GenRefreshToken 生成Refresh Token（单独函数）
func GenRefreshToken(userID uint) (string, error) {
	claims := MyClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(
				time.Duration(global.Config.Jwt.RefreshExpire) * time.Hour)),
			Issuer: global.Config.Jwt.Issuer + "_refresh", // 建议区分issuer
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(global.Config.Jwt.RefreshSecret))
}

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
