package security

import (
	"brodo-demo/service/security"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtTokenManager struct {
}

func NewJwtTokenManager() security.AuthenticationTokenManager {
	return &JwtTokenManager{}
}

func getSecretKey() []byte {
	return []byte(os.Getenv("JWT_SECRET_KEY"))
}

type jwtClaims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

func (tokenManager *JwtTokenManager) CreateAccessToken(userId int) (string, error) {
	claims := jwtClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(getSecretKey())
}

func (tokenManager *JwtTokenManager) VerifyAccessToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &jwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		return getSecretKey(), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*jwtClaims); ok && token.Valid {
		return claims.UserId, nil
	} else {
		return 0, errors.New("invalid token")
	}
}
