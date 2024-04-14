package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/sema0205/avito-backend-assignment-2024/pkg/auth"
	"time"
)

type TokenManager struct {
	accessTokenTTL time.Duration
	signPrivateKey string
}

func NewTokenManager(signPrivateKey string, accessTokenTTL time.Duration) (*TokenManager, error) {
	if signPrivateKey == "" {
		return nil, errors.New("empty sign private key")
	}

	return &TokenManager{
		accessTokenTTL: accessTokenTTL,
		signPrivateKey: signPrivateKey,
	}, nil
}

func (m *TokenManager) NewAdminJWT(adminId int) (string, error) {
	return m.newJWT(auth.Admin, adminId)
}

func (m *TokenManager) NewUserJWT(userId int) (string, error) {
	return m.newJWT(auth.User, userId)
}

func (m *TokenManager) newJWT(role auth.Role, id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Role: role,
		Id:   id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(m.accessTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	return token.SignedString([]byte(m.signPrivateKey))
}

func (m *TokenManager) Parse(accessToken string) (auth.TokenClaims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signPrivateKey), nil
	})

	if err != nil || !token.Valid {
		return auth.TokenClaims{}, fmt.Errorf("invalid or expired token")
	}

	return auth.TokenClaims{
		Role: claims.Role,
		Id:   claims.Id,
	}, nil
}
