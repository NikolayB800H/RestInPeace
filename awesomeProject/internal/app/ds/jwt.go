package ds

import (
	"awesomeProject/internal/app/role"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	UserUUID string `json:"user_uuid"`
	Role     role.Role
	Login    string `json:"login"`
}
