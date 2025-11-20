package jwt

import (
	"fmt"
	"rohmatext/ore-note/internal/usecase"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MapClaims struct {
	UserId uint `json:"user_id"`
	jwt.RegisteredClaims
}

type JWTService struct {
	secret string
}

func NewJWTService(secret string) usecase.TokenService {
	return &JWTService{secret: secret}
}

func (j *JWTService) GenerateToken(userId uint) (string, error) {
	claims := &MapClaims{
		userId,
		jwt.RegisteredClaims{
			Subject:   fmt.Sprintf("user:%d", userId),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 15)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secret))
}
