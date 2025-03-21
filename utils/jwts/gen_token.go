package jwts

import (
	"FastFiber/global"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

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
