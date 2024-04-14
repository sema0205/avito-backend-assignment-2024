package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/sema0205/avito-backend-assignment-2024/pkg/auth"
)

type Claims struct {
	Role auth.Role `json:"role"`
	Id   int       `json:"id"`
	jwt.StandardClaims
}
