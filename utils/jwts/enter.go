package jwts

import (
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	UserID   uint   `json:"user_id"`
	Nickname string `json:"nickname"`
	RoleID   string `json:"role_id"`
	jwt.RegisteredClaims
}
